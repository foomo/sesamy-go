package client

import (
	"fmt"
	"io"
	"net/http"

	"github.com/foomo/sesamy-go/pkg/encoding/gtag"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type (
	GTag struct {
		l               *zap.Logger
		path            string
		host            string
		cookies         []string
		trackingID      string
		protocolVersion string
		httpClient      *http.Client
		middlewares     []GTagMiddleware
	}
	GTagOption     func(*GTag)
	GTagHandler    func(r *http.Request, payload *gtag.Payload) error
	GTagMiddleware func(next GTagHandler) GTagHandler
)

// ------------------------------------------------------------------------------------------------
// ~ Options
// ------------------------------------------------------------------------------------------------

func GTagWithHTTPClient(v *http.Client) GTagOption {
	return func(o *GTag) {
		o.httpClient = v
	}
}

func GTagWithPath(v string) GTagOption {
	return func(o *GTag) {
		o.path = v
	}
}

func GTagWithCookies(v ...string) GTagOption {
	return func(o *GTag) {
		o.cookies = append(o.cookies, v...)
	}
}

func GTagWithMiddlewares(v ...GTagMiddleware) GTagOption {
	return func(o *GTag) {
		o.middlewares = append(o.middlewares, v...)
	}
}

// ------------------------------------------------------------------------------------------------
// ~ Constructor
// ------------------------------------------------------------------------------------------------

func NewGTag(l *zap.Logger, host, trackingID string, opts ...GTagOption) *GTag {
	inst := &GTag{
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
		GTagMiddlewareRichsstsse,
		GTagMiddlewareTrackingID(inst.trackingID),
		GTagMiddlewarProtocolVersion("2"),
		GTagMiddlewarIsDebug,
		GTagMiddlewarClientID,
		GTagMiddlewarSessionID(inst.trackingID),
	)
	return inst
}

// ------------------------------------------------------------------------------------------------
// ~ Getter
// ------------------------------------------------------------------------------------------------

func (c *GTag) HTTPClient() *http.Client {
	return c.httpClient
}

// ------------------------------------------------------------------------------------------------
// ~ Public methods
// ------------------------------------------------------------------------------------------------

func (c *GTag) Send(r *http.Request, payload *gtag.Payload) error {
	next := c.SendRaw
	for _, middleware := range c.middlewares {
		next = middleware(next)
	}
	return next(r, payload)
}

func (c *GTag) SendRaw(r *http.Request, payload *gtag.Payload) error {
	values, body, err := gtag.Encode(payload)
	if err != nil {
		return errors.Wrap(err, "failed to encode payload")
	}

	req, err := http.NewRequestWithContext(
		r.Context(),
		http.MethodPost,
		fmt.Sprintf("%s%s?%s", c.host, c.path, gtag.EncodeValues(values)),
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
	defer resp.Body.Close()

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
