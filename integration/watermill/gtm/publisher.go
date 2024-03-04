package gtm

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	mpv2 "github.com/foomo/sesamy/measurementprotocol/v2"
	"github.com/pkg/errors"
)

var (
	ErrErrorResponse   = errors.New("server responded with error status")
	ErrPublisherClosed = errors.New("publisher is closed")
)

type (
	Publisher struct {
		//l                  watermill.LoggerAdapter
		url                string
		client             *http.Client
		marshalMessageFunc MarshalMessageFunc
		closed             bool
	}
	PublisherOption func(*Publisher)
)

// ------------------------------------------------------------------------------------------------
// ~ Constructor
// ------------------------------------------------------------------------------------------------

func NewPublisher(url string, opts ...PublisherOption) *Publisher {
	inst := &Publisher{
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

func PublisherWithMarshalMessageFunc(v MarshalMessageFunc) PublisherOption {
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

		values, body, err := mpv2.Marshal(event)
		if err != nil {
			return err
		}

		u := p.url + "?" + values.Encode() + "&richsstsse"

		req, err := http.NewRequestWithContext(msg.Context(), "POST", u, body)

		for s, s2 := range msg.Metadata {
			req.Header.Set(s, s2)
		}

		//req, err := p.marshalMessageFunc(topic, msg)
		if err != nil {
			return errors.Wrapf(err, "cannot marshal message %s", msg.UUID)
		}

		logFields := watermill.LogFields{
			"uuid":     msg.UUID,
			"provider": ProviderName,
		}

		//p.l.Trace("Publishing message", logFields)

		resp, err := p.client.Do(req)
		if err != nil {
			return errors.Wrapf(err, "publishing message %s failed", msg.UUID)
		}

		if err = p.handleResponseBody(resp, logFields); err != nil {
			return err
		}

		if resp.StatusCode >= http.StatusBadRequest {
			return errors.Wrap(ErrErrorResponse, resp.Status)
		}

		//p.l.Trace("Message published", logFields)
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

func (p *Publisher) handleResponseBody(resp *http.Response, logFields watermill.LogFields) error {
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusBadRequest {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "could not read response body")
	}

	logFields = logFields.Add(watermill.LogFields{
		"http_status":   resp.StatusCode,
		"http_response": string(body),
	})
	//p.l.Info("Server responded with error", logFields)

	return nil
}

// MarshalMessageFunc transforms the message into a HTTP request to be sent to the specified url.
type MarshalMessageFunc func(url string, msg *message.Message) (*http.Request, error)

// DefaultMarshalMessageFunc transforms the message into a HTTP POST request.
// It encodes the UUID and Metadata in request headers.
func DefaultMarshalMessageFunc(url string, msg *message.Message) (*http.Request, error) {

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(msg.Payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set(HeaderUUID, msg.UUID)

	return req, nil
}
