package mpv2

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type (
	Publisher struct {
		l           *zap.Logger
		host        string
		path        string
		httpClient  *http.Client
		middlewares []PublisherMiddleware
		closed      bool
	}
	PublisherOption     func(*Publisher)
	PublisherHandler    func(l *zap.Logger, msg *message.Message) error
	PublisherMiddleware func(next PublisherHandler) PublisherHandler
)

// ------------------------------------------------------------------------------------------------
// ~ Constructor
// ------------------------------------------------------------------------------------------------

func NewPublisher(l *zap.Logger, host string, opts ...PublisherOption) *Publisher {
	inst := &Publisher{
		l:          l,
		host:       host,
		path:       "/mp/collect",
		httpClient: http.DefaultClient,
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

func PublisherWithHTTPClient(v *http.Client) PublisherOption {
	return func(o *Publisher) {
		o.httpClient = v
	}
}

func PublisherWithMiddlewares(v ...PublisherMiddleware) PublisherOption {
	return func(o *Publisher) {
		o.middlewares = append(o.middlewares, v...)
	}
}

// ------------------------------------------------------------------------------------------------
// ~ Getter
// ------------------------------------------------------------------------------------------------

func (p *Publisher) HTTPClient() *http.Client {
	return p.httpClient
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
	req, err := http.NewRequestWithContext(msg.Context(), http.MethodPost, fmt.Sprintf("%s%s", p.host, p.path), bytes.NewReader(msg.Payload))
	if err != nil {
		return errors.Wrap(err, "failed to create request")
	}

	for s, s2 := range msg.Metadata {
		if s == "Cookie" {
			for _, s3 := range strings.Split(s2, "; ") {
				val := strings.Split(s3, "=")
				req.AddCookie(&http.Cookie{
					Name:  val[0],
					Value: strings.Join(val[1:], "="),
				})
			}
		} else {
			req.Header.Set(s, s2)
		}
	}

	if err := func() error {
		resp, err := p.httpClient.Do(req)
		if err != nil {
			return errors.Wrapf(err, "failed to publish message: %s", msg.UUID)
		}
		defer resp.Body.Close()

		l = l.With(zap.Int("http_status_code", resp.StatusCode))

		if resp.StatusCode >= http.StatusBadRequest {
			if body, err := io.ReadAll(resp.Body); err == nil {
				l = l.With(zap.String("http_response", string(body)))
			}
			l.Warn("server responded with error")
			return errors.Wrap(ErrErrorResponse, resp.Status)
		}

		l.Debug("message published")

		return nil
	}(); err != nil {
		return err
	}

	return nil
}
