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
	// GTag sends gtag.js payloads to a server-side tagging endpoint.
	// Create one with [NewGTag].
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
	// GTagOption configures a [GTag] client.
	GTagOption func(*GTag)
	// GTagHandler sends payload in the context of the incoming request r.
	GTagHandler func(r *http.Request, payload *gtag.Payload) error
	// GTagMiddleware wraps a [GTagHandler], typically to enrich the payload.
	GTagMiddleware func(next GTagHandler) GTagHandler
)

// ------------------------------------------------------------------------------------------------
// ~ Options
// ------------------------------------------------------------------------------------------------

// GTagWithHTTPClient sets the HTTP client used for upstream requests.
// The default is [http.DefaultClient].
func GTagWithHTTPClient(v *http.Client) GTagOption {
	return func(o *GTag) {
		o.httpClient = v
	}
}

// GTagWithPath sets the collect path appended to the host.
// The default is "/g/collect".
func GTagWithPath(v string) GTagOption {
	return func(o *GTag) {
		o.path = v
	}
}

// GTagWithCookies appends names of cookies forwarded from the incoming request.
// The defaults are gtm_auth, gtm_debug and gtm_preview.
func GTagWithCookies(v ...string) GTagOption {
	return func(o *GTag) {
		o.cookies = append(o.cookies, v...)
	}
}

// GTagWithMiddlewares appends middlewares.
// Middlewares wrap in list order, so the last one runs first (outermost).
func GTagWithMiddlewares(v ...GTagMiddleware) GTagOption {
	return func(o *GTag) {
		o.middlewares = append(o.middlewares, v...)
	}
}

// ------------------------------------------------------------------------------------------------
// ~ Constructor
// ------------------------------------------------------------------------------------------------

// NewGTag returns a [GTag] client that sends to host with the GA4 measurement ID
// trackingID. Custom middlewares from [GTagWithMiddlewares] run inside the
// built-in ones, which set richsstsse, tracking ID, protocol version, debug flag,
// client ID and session ID.
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

// HTTPClient returns the HTTP client used for upstream requests.
func (c *GTag) HTTPClient() *http.Client {
	return c.httpClient
}

// ------------------------------------------------------------------------------------------------
// ~ Public methods
// ------------------------------------------------------------------------------------------------

// Send runs payload through the middlewares and sends it with [GTag.SendRaw].
// r is the incoming request whose context, headers and cookies are used.
func (c *GTag) Send(r *http.Request, payload *gtag.Payload) error {
	next := c.SendRaw
	for _, middleware := range c.middlewares {
		next = middleware(next)
	}

	return next(r, payload)
}

// SendRaw encodes payload and posts it to the tagging endpoint without running
// middlewares. It copies the headers of r except Cookie and Authorization and
// forwards only the configured cookies. It returns an error if the upstream
// response status is not 200 OK.
func (c *GTag) SendRaw(r *http.Request, payload *gtag.Payload) error {
	query, body, err := gtag.Encode(payload)
	if err != nil {
		return errors.Wrap(err, "failed to encode payload")
	}

	req, err := http.NewRequestWithContext(
		r.Context(),
		http.MethodPost,
		fmt.Sprintf("%s%s?%s", c.host, c.path, query),
		body,
	)
	if err != nil {
		return errors.Wrap(err, "failed to create request")
	}

	// copy headers without credentials, cookies are forwarded explicitly below
	req.Header = r.Header.Clone()
	req.Header.Del("Cookie")
	req.Header.Del("Authorization")

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
