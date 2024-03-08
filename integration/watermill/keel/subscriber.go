package keel

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
	keellog "github.com/foomo/keel/net/http/log"
	keelhttputils "github.com/foomo/keel/utils/net/http"
	mp "github.com/foomo/sesamy-go/measurementprotocol"
	mpv2 "github.com/foomo/sesamy-go/measurementprotocol/v2"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var (
	ErrMissingEventName = errors.New("missing event name")
	ErrContextCanceled  = errors.New("request stopped without ACK received")
	ErrMessageNacked    = errors.New("message nacked")
	ErrClosed           = errors.New("subscriber already closed")
)

type (
	Subscriber struct {
		l           *zap.Logger
		uuidFunc    func() string
		messages    chan *message.Message
		middlewares []SubscriberMiddleware
		closed      bool
	}
	SubscriberOption     func(*Subscriber)
	SubscriberHandler    func(l *zap.Logger, r *http.Request, event *mpv2.Event) error
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

func SubscriberWithMiddlewares(v ...SubscriberMiddleware) SubscriberOption {
	return func(o *Subscriber) {
		o.middlewares = append(o.middlewares, v...)
	}
}

func SubscriberWithLogger(fields ...zap.Field) SubscriberOption {
	return func(o *Subscriber) {
		o.middlewares = append(o.middlewares, func(next SubscriberHandler) SubscriberHandler {
			return func(l *zap.Logger, r *http.Request, event *mpv2.Event) error {
				fields := append(fields, zap.String("event_name", mp.GetDefault(event.EventName, "-").String()))
				if labeler, ok := keellog.LabelerFromRequest(r); ok {
					labeler.Add(fields...)
				}
				return next(l.With(fields...), r, event)
			}
		})
	}
}

func SubscriberWithRequireEventName() SubscriberOption {
	return func(o *Subscriber) {
		o.middlewares = append(o.middlewares, func(next SubscriberHandler) SubscriberHandler {
			return func(l *zap.Logger, r *http.Request, event *mpv2.Event) error {
				if event.EventName == nil {
					return ErrMissingEventName
				}
				return next(l, r, event)
			}
		})
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
		keelhttputils.ServerError(s.l, w, r, http.StatusMethodNotAllowed, errors.New(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}

	// unmarshal event
	var event *mpv2.Event
	if err := mpv2.Decode(values, &event); err != nil {
		keelhttputils.InternalServerError(s.l, w, r, errors.Wrap(err, "failed to marshal url values"))
		return
	}

	// compose middlewares
	next := s.handle
	for _, middleware := range s.middlewares {
		next = middleware(next)
	}

	// run handler
	if err := next(s.l, r, event); err != nil {
		keelhttputils.InternalServerError(s.l, w, r, err)
		return
	}
}

func (s *Subscriber) handle(l *zap.Logger, r *http.Request, event *mpv2.Event) error {
	// marshal message payload
	payload, err := json.Marshal(event)
	if err != nil {
		return errors.Wrap(err, "failed to marshal payload")
	}

	msg := message.NewMessage(s.uuidFunc(), payload)
	l = l.With(zap.String("message_id", msg.UUID))
	if labeler, ok := keellog.LabelerFromRequest(r); ok {
		labeler.Add(zap.String("message_id", msg.UUID))
	}

	if event.EventName != nil {
		msg.Metadata.Set(MetadataEventName, mp.Get(event.EventName).String())
	}

	// TODO filter headers?
	for name, headers := range r.Header {
		msg.Metadata.Set(name, strings.Join(headers, ","))
	}
	//
	// if cookies := r.Cookies(); len(cookies) > 0 {
	// 	values := make([]string, len(cookies))
	// 	for i, cookie := range r.Cookies() {
	// 		values[i] = cookie.String()
	// 	}
	// 	msg.Metadata.Set("Cookie", strings.Join(values, "; "))
	// }

	for k, v := range msg.Metadata {
		l = l.With(zap.String(k, v))
	}

	// TODO different context?
	ctx, cancelCtx := context.WithCancel(r.Context())
	msg.SetContext(ctx)
	defer cancelCtx()

	// send message
	s.messages <- msg

	// wait for ACK
	select {
	case <-msg.Acked():
		l.Info("message acked")
		return nil
	case <-msg.Nacked():
		l.Info("message nacked")
		return ErrMessageNacked
	case <-r.Context().Done():
		l.Info("message cancled")
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
