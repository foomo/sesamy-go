package mpv2

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/foomo/sesamy-go/pkg/encoding/mpv2"
	mpv2http "github.com/foomo/sesamy-go/pkg/http/mpv2"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type (
	Subscriber struct {
		l           *zap.Logger
		uuidFunc    func() string
		messages    chan *message.Message
		messageFunc func(l *zap.Logger, r *http.Request, msg *message.Message) error
		middlewares []mpv2http.Middleware
		closed      bool
	}
	SubscriberOption func(*Subscriber)
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

func SubscriberWithMiddlewares(v ...mpv2http.Middleware) SubscriberOption {
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
	// retrieve payload
	payload := mpv2http.Handler(w, r)

	// compose middlewares
	next := s.handle
	for _, middleware := range s.middlewares {
		next = middleware(next)
	}

	// run handler
	if err := next(s.l, w, r, payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Subscriber) handle(l *zap.Logger, w http.ResponseWriter, r *http.Request, payload *mpv2.Payload[any]) error {
	// marshal message payload
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return errors.Wrap(err, "failed to marshal payload")
	}

	msg := message.NewMessage(s.uuidFunc(), jsonPayload)
	l = l.With(zap.String("message_id", msg.UUID))
	msg.SetContext(context.WithoutCancel(r.Context()))

	// store query
	msg.Metadata.Set(MetadataRequestQuery, r.URL.RawQuery)

	// store header
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
