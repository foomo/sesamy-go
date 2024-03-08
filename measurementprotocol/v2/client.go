package v2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type (
	Client struct {
		l               *zap.Logger
		url             string
		cookies         []string
		richsstsse      string
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

func NewClient(l *zap.Logger, url, trackingID string, opts ...ClientOption) *Client {
	inst := &Client{
		l:               l,
		url:             url,
		cookies:         []string{"gtm_auth", "gtm_debug", "gtm_preview"},
		trackingID:      trackingID,
		protocolVersion: "2",
		httpClient:      http.DefaultClient,
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

func (c *Client) Send(r *http.Request, event *Event) error {
	yes := "1"

	// set default values
	event.TrackingID = &c.trackingID
	event.Richsstsse = &c.richsstsse
	event.ProtocolVersion = &c.protocolVersion

	event.IgnoreReferrer = &yes

	{ // set referrer parameter
		if referrer, err := url.Parse(r.Referer()); err != nil {
			c.l.With(zap.Error(err)).Warn("failed to parse referrer")
		} else {
			event.DocumentLocation = &referrer.Path
			event.DocumentHostname = &referrer.Host
		}
	}

	{ // TODO check
		if value, _ := r.Cookie("gtm_debug"); value != nil {
			event.IsDebug = &yes
		}
	}

	{ // set client id
		if value, _ := r.Cookie("_ga"); value != nil {
			clientID := strings.TrimPrefix(value.Value, "GA1.1.")
			event.ClientID = &clientID
		}
	}

	next := c.SendRaw
	for _, middleware := range c.middlewares {
		next = middleware(next)
	}

	return next(r, event)
}

func (c *Client) SendRaw(r *http.Request, event *Event) error {
	values, body, err := Encode(event)
	if err != nil {
		return errors.Wrap(err, "failed to marshall event")
	}

	req, err := http.NewRequestWithContext(
		r.Context(),
		http.MethodPost,
		fmt.Sprintf("%s?%s", c.url, EncodeValues(values)),
		body,
	)
	if err != nil {
		return errors.Wrap(err, "failed to create request")
	}

	// copy headers
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
