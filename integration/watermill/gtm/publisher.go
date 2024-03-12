package gtm

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/foomo/keel/log"
	mpv2 "github.com/foomo/sesamy-go/measurementprotocol/v2"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var (
	ErrErrorResponse   = errors.New("server responded with error status")
	ErrPublisherClosed = errors.New("publisher is closed")
)

type (
	Publisher struct {
		l                  *zap.Logger
		url                string
		client             *http.Client
		marshalMessageFunc PublisherMarshalMessageFunc
		closed             bool
	}
	PublisherOption func(*Publisher)
	// PublisherMarshalMessageFunc transforms the message into a HTTP request to be sent to the specified url.
	PublisherMarshalMessageFunc func(url string, msg *message.Message) (*http.Request, error)
)

// ------------------------------------------------------------------------------------------------
// ~ Constructor
// ------------------------------------------------------------------------------------------------

func NewPublisher(l *zap.Logger, url string, opts ...PublisherOption) *Publisher {
	inst := &Publisher{
		l:      l,
		url:    url,
		client: http.DefaultClient,
	}
	for _, opt := range opts {
		opt(inst)
	}
	return inst
}

// ------------------------------------------------------------------------------------------------
// ~ Options
// ------------------------------------------------------------------------------------------------

func PublisherWithClient(v *http.Client) PublisherOption {
	return func(o *Publisher) {
		o.client = v
	}
}

func PublisherWithMarshalMessageFunc(v PublisherMarshalMessageFunc) PublisherOption {
	return func(o *Publisher) {
		o.marshalMessageFunc = v
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
		var event *mpv2.Event
		if err := json.Unmarshal(msg.Payload, &event); err != nil {
			return err
		}

		values, body, err := mpv2.Encode(event)
		if err != nil {
			return err
		}

		req, err := http.NewRequestWithContext(msg.Context(), http.MethodPost, fmt.Sprintf("%s?%s", p.url, mpv2.EncodeValues(values)), body)
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

		l := log.WithHTTPRequestOut(p.l, req).With(
			zap.String("message_id", msg.UUID),
		)

		resp, err := p.client.Do(req)
		if err != nil {
			return errors.Wrapf(err, "failed to publish message: %s", msg.UUID)
		}
		defer resp.Body.Close()

		l = l.With(log.FHTTPStatusCode(resp.StatusCode))

		if resp.StatusCode >= http.StatusBadRequest {
			if body, err := io.ReadAll(resp.Body); err == nil {
				l = l.With(zap.String("http_response", string(body)))
			}
			l.Info("server responded with error")
			return errors.Wrap(ErrErrorResponse, resp.Status)
		}

		l.Debug("message published")
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
