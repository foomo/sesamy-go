package mpv2

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/foomo/sesamy-go/pkg/encoding/mpv2"
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
	SubscriberHandler    func(l *zap.Logger, r *http.Request, payload *mpv2.Payload[any]) error
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

// ------------------------------------------------------------------------------------------------
// ~ Public methods
// ------------------------------------------------------------------------------------------------

func (s *Subscriber) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// read request body
	var payload *mpv2.Payload[any]
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// validate required fields
	if len(payload.Events) == 0 {
		http.Error(w, "missing events", http.StatusBadRequest)
		return
	}
	for _, event := range payload.Events {
		if event.Name == "" {
			http.Error(w, "missing event name", http.StatusBadRequest)
			return
		}
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

func (s *Subscriber) handle(l *zap.Logger, r *http.Request, payload *mpv2.Payload[any]) error {
	// marshal message payload
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return errors.Wrap(err, "failed to marshal payload")
	}

	msg := message.NewMessage(s.uuidFunc(), jsonPayload)
	l = l.With(zap.String("message_id", msg.UUID))
	msg.SetContext(context.WithoutCancel(r.Context()))

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
