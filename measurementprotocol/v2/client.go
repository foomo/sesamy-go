package v2

import (
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type (
	Client struct {
		l               *zap.Logger
		path            string
		host            string
		cookies         []string
		trackingID      string
		measurementID   string
		protocolVersion string
		httpClient      *http.Client
		middlewares     []ClientMiddleware
	}
	ClientOption     func(*Client)
	ClientHandler    func(r *http.Request, event *Event) error
	ClientMiddleware func(next ClientHandler) ClientHandler
)

// ------------------------------------------------------------------------------------------------
// ~ Options
// ------------------------------------------------------------------------------------------------

func ClientWithHTTPClient(v *http.Client) ClientOption {
	return func(o *Client) {
		o.httpClient = v
	}
}

func ClientWithPath(v string) ClientOption {
	return func(o *Client) {
		o.path = v
	}
}

func ClientWithCookies(v ...string) ClientOption {
	return func(o *Client) {
		o.cookies = append(o.cookies, v...)
	}
}

func ClientWithMiddlewares(v ...ClientMiddleware) ClientOption {
	return func(o *Client) {
		o.middlewares = append(o.middlewares, v...)
	}
}

// ------------------------------------------------------------------------------------------------
// ~ Constructor
// ------------------------------------------------------------------------------------------------

func NewClient(l *zap.Logger, host, trackingID string, opts ...ClientOption) *Client {
	inst := &Client{
		l:               l,
		host:            host,
		path:            "/g/collect",
		cookies:         []string{"gtm_auth", "gtm_debug", "gtm_preview"},
		trackingID:      trackingID,
		protocolVersion: "2",
		httpClient:      http.DefaultClient,
	}
	for _, opt := range opts {
		opt(inst)
	}
	inst.middlewares = append(inst.middlewares,
		MiddlewareRichsstsse,
		MiddlewareTrackingID(inst.trackingID),
		MiddlewarIgnoreReferrer("1"),
		MiddlewarProtocolVersion("2"),
		MiddlewarDebug,
		MiddlewarClientID,
		MiddlewarDocument,
	)
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

func (c *Client) Send(r *http.Request, event Marshler) error {
	e, err := event.MarshalMPv2()
	if err != nil {
		return err
	}
	return c.SendEvent(r, e)
}

func (c *Client) SendEvent(r *http.Request, event *Event) error {
	next := c.SendRawEvent
	for _, middleware := range c.middlewares {
		next = middleware(next)
	}
	return next(r, event)
}

func (c *Client) SendRawEvent(r *http.Request, event *Event) error {
	values, body, err := Encode(event)
	if err != nil {
		return errors.Wrap(err, "failed to encode event")
	}

	req, err := http.NewRequestWithContext(
		r.Context(),
		http.MethodPost,
		fmt.Sprintf("%s%s?%s", c.host, c.path, EncodeValues(values)),
		body,
	)
	if err != nil {
		return errors.Wrap(err, "failed to create request")
	}

	// TODO valiate: copy headers
	req.Header = r.Header.Clone()

	// forward cookies
	for _, cookie := range c.cookies {
		if value, _ := r.Cookie(cookie); value != nil {
			req.AddCookie(value)
		}
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
