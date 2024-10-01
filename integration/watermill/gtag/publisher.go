package gtag

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/foomo/sesamy-go/pkg/encoding/gtag"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var (
	ErrErrorResponse   = errors.New("server responded with error status")
	ErrPublisherClosed = errors.New("publisher is closed")
)

type (
	Publisher struct {
		l               *zap.Logger
		host            string
		path            string
		client          *http.Client
		closed          bool
		middlewares     []PublisherMiddleware
		maxResponseCode int
	}
	PublisherOption     func(*Publisher)
	PublisherHandler    func(l *zap.Logger, msg *message.Message) error
	PublisherMiddleware func(next PublisherHandler) PublisherHandler
	// PublisherMarshalMessageFunc transforms the message into a HTTP request to be sent to the specified url.
	PublisherMarshalMessageFunc func(url string, msg *message.Message) (*http.Request, error)
)

// ------------------------------------------------------------------------------------------------
// ~ Constructor
// ------------------------------------------------------------------------------------------------

func NewPublisher(l *zap.Logger, host string, opts ...PublisherOption) *Publisher {
	inst := &Publisher{
		l:               l,
		host:            host,
		path:            "/g/collect",
		client:          http.DefaultClient,
		maxResponseCode: http.StatusBadRequest,
	}
	for _, opt := range opts {
		opt(inst)
	}
	return inst
}

// ------------------------------------------------------------------------------------------------
// ~ Options
// ------------------------------------------------------------------------------------------------

func PublisherWithPath(v string) PublisherOption {
	return func(o *Publisher) {
		o.path = v
	}
}

func PublisherWithClient(v *http.Client) PublisherOption {
	return func(o *Publisher) {
		o.client = v
	}
}

func PublisherWithMiddlewares(v ...PublisherMiddleware) PublisherOption {
	return func(o *Publisher) {
		o.middlewares = append(o.middlewares, v...)
	}
}

func PublisherWithMaxResponseCode(v int) PublisherOption {
	return func(o *Publisher) {
		o.maxResponseCode = v
	}
}

// ------------------------------------------------------------------------------------------------
// ~ Getter
// ------------------------------------------------------------------------------------------------

func (p *Publisher) Client() *http.Client {
	return p.client
}

// ------------------------------------------------------------------------------------------------
// ~ Public methods
// ------------------------------------------------------------------------------------------------

func (p *Publisher) Publish(topic string, messages ...*message.Message) error {
	if p.closed {
		return ErrPublisherClosed
	}

	for _, msg := range messages {
		// compose middlewares
		next := p.handle
		for _, middleware := range p.middlewares {
			next = middleware(next)
		}

		// run handler
		if err := next(p.l.With(
			zap.String("message_id", msg.UUID),
		), msg); err != nil {
			return err
		}
	}

	return nil
}

func (p *Publisher) Close() error {
	if p.closed {
		return nil
	}

	p.closed = true
	return nil
}

// ------------------------------------------------------------------------------------------------
// ~ Private methods
// ------------------------------------------------------------------------------------------------

func (p *Publisher) handle(l *zap.Logger, msg *message.Message) error {
	var event *gtag.Payload
	if err := json.Unmarshal(msg.Payload, &event); err != nil {
		return err
	}

	values, body, err := gtag.Encode(event)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(msg.Context(), http.MethodPost, fmt.Sprintf("%s%s?%s", p.host, p.path, gtag.EncodeValues(values)), body)
	if err != nil {
		return errors.Wrap(err, "failed to create request")
	}

	for s, s2 := range msg.Metadata {
		req.Header.Set(s, s2)
	}

	if err := func() error {
		resp, err := p.client.Do(req)
		if err != nil {
			return errors.Wrapf(err, "failed to publish message: %s", msg.UUID)
		}
		defer resp.Body.Close()

		l = l.With(zap.Int("http_status_code", resp.StatusCode))

		if p.maxResponseCode > 0 && resp.StatusCode >= p.maxResponseCode {
			if body, err := io.ReadAll(resp.Body); err == nil {
				l = l.With(zap.String("http_response", string(body)))
			}
			return errors.Wrap(ErrErrorResponse, resp.Status)
		}

		return nil
	}(); err != nil {
		return err
	}

	return nil
}
