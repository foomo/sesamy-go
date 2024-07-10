package gtag

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/foomo/sesamy-go/pkg/encoding/gtag"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type (
	Subscriber struct {
		l           *zap.Logger
		uuidFunc    func() string
		messages    chan *message.Message
		messageFunc func(l *zap.Logger, r *http.Request, msg *message.Message) error
		middlewares []SubscriberMiddleware
		closed      bool
	}
	SubscriberOption     func(*Subscriber)
	SubscriberHandler    func(l *zap.Logger, r *http.Request, payload *gtag.Payload) error
	SubscriberMiddleware func(next SubscriberHandler) SubscriberHandler
)

// ------------------------------------------------------------------------------------------------
// ~ Options
// ------------------------------------------------------------------------------------------------

func SubscriberWithUUIDFunc(v func() string) SubscriberOption {
	return func(o *Subscriber) {
		o.uuidFunc = v
	}
}

func SubscriberWithMessageFunc(v func(l *zap.Logger, r *http.Request, msg *message.Message) error) SubscriberOption {
	return func(o *Subscriber) {
		o.messageFunc = v
	}
}

func SubscriberWithMiddlewares(v ...SubscriberMiddleware) SubscriberOption {
	return func(o *Subscriber) {
		o.middlewares = append(o.middlewares, v...)
	}
}

// ------------------------------------------------------------------------------------------------
// ~ Constructor
// ------------------------------------------------------------------------------------------------

func NewSubscriber(l *zap.Logger, opts ...SubscriberOption) *Subscriber {
	inst := &Subscriber{
		l:        l,
		uuidFunc: watermill.NewUUID,
		messages: make(chan *message.Message),
	}
	for _, opt := range opts {
		opt(inst)
	}
	return inst
}

func (s *Subscriber) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var values url.Values

	switch r.Method {
	case http.MethodGet:
		values = r.URL.Query()
	case http.MethodPost:
		values = r.URL.Query()

		// read request body
		out, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to read body: %s", err.Error()), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		// append request body to query
		if len(out) > 0 {
			v, err := url.ParseQuery(string(out))
			if err != nil {
				http.Error(w, fmt.Sprintf("failed to parse extended url: %s", err.Error()), http.StatusInternalServerError)
				return
			}
			for s2, i := range v {
				values.Set(s2, i[0])
			}
		} else {
			values = r.URL.Query()
		}
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// unmarshal event
	var payload *gtag.Payload
	if err := gtag.Decode(values, &payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// validate
	if payload.EventName == nil || payload.EventName.String() == "" {
		http.Error(w, "missing event name", http.StatusBadRequest)
		return
	}

	// compose middlewares
	next := s.handle
	for _, middleware := range s.middlewares {
		next = middleware(next)
	}

	// run handler
	if err := next(s.l, r, payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Subscriber) handle(l *zap.Logger, r *http.Request, payload *gtag.Payload) error {
	// marshal message payload
	data, err := json.Marshal(payload)
	if err != nil {
		return errors.Wrap(err, "failed to marshal payload")
	}

	msg := message.NewMessage(s.uuidFunc(), data)
	l = l.With(zap.String("message_id", msg.UUID))
	msg.SetContext(context.WithoutCancel(r.Context()))

	if payload.EventName != nil {
		msg.Metadata.Set(MetadataEventName, gtag.Get(payload.EventName).String())
	}

	for name, headers := range r.Header {
		msg.Metadata.Set(name, strings.Join(headers, ","))
	}

	if s.messageFunc != nil {
		if err := s.messageFunc(l, r, msg); err != nil {
			return err
		}
	}

	for k, v := range msg.Metadata {
		l = l.With(zap.String(k, v))
	}

	// send message
	s.messages <- msg

	// wait for ACK
	select {
	case <-msg.Acked():
		l.Debug("message acked")
		return nil
	case <-msg.Nacked():
		l.Debug("message nacked")
		return ErrMessageNacked
	case <-r.Context().Done():
		l.Debug("message canceled")
		return ErrContextCanceled
	}
}

func (s *Subscriber) Subscribe(ctx context.Context, topic string) (<-chan *message.Message, error) {
	return s.messages, nil
}

// Close closes all subscriptions with their output channels and flush offsets etc. when needed.
func (s *Subscriber) Close() error {
	if s.closed {
		return ErrClosed
	}
	s.closed = true

	close(s.messages)
	return nil
}
