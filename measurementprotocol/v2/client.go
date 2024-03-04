package v2

import (
	"context"
	"io"
	"net/http"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type (
	Client struct {
		l          *zap.Logger
		url        string
		httpClient *http.Client
	}
	ClientOption func(*Client)
)

// ------------------------------------------------------------------------------------------------
// ~ Options
// ------------------------------------------------------------------------------------------------

func ClientWithHTTPClient(v *http.Client) ClientOption {
	return func(o *Client) {
		o.httpClient = v
	}
}

// ------------------------------------------------------------------------------------------------
// ~ Constructor
// ------------------------------------------------------------------------------------------------

func NewClient(l *zap.Logger, url string, opts ...ClientOption) *Client {
	inst := &Client{
		l:          l,
		url:        url,
		httpClient: http.DefaultClient,
	}
	for _, opt := range opts {
		opt(inst)
	}
	return inst
}

// ------------------------------------------------------------------------------------------------
// ~ Getter
// ------------------------------------------------------------------------------------------------

func (c *Client) HTTPClient() *http.Client {
	return c.httpClient
}

// ------------------------------------------------------------------------------------------------
// ~ Public methods
// ------------------------------------------------------------------------------------------------

func (c *Client) Send(ctx context.Context, event *Event) error {

	values, body, err := Marshal(event)
	if err != nil {
		return errors.Wrap(err, "failed to marshall event")
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.url+"?"+values.Encode(),
		body,
	)
	if err != nil {
		return errors.Wrap(err, "failed to create request")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to send request")
	}

	if resp.StatusCode != http.StatusOK {
		var body string
		if out, err := io.ReadAll(resp.Body); err != nil {
			c.l.With(zap.Error(err)).Warn(err.Error())
		} else {
			body = string(out)
		}
		return errors.Errorf("unexpected response status: %d (%s)", resp.StatusCode, body)
	}

	return nil
}

