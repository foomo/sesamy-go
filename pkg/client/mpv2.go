package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/foomo/sesamy-go/pkg/encoding/mpv2"
	"github.com/foomo/sesamy-go/pkg/sesamy"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type (
	// MPv2 sends Measurement Protocol v2 payloads to a server-side tagging
	// endpoint. Create one with [NewMPv2].
	MPv2 struct {
		l       *zap.Logger
		path    string
		host    string
		cookies []string
		// To create a new secret, navigate in the Google Analytics UI to:
		// Admin > Data Streams > choose your stream > Measurement Protocol > Create
		apiSecret string
		// Measurement ID. The identifier for a Data Stream. Found in the Google Analytics UI under:
		// Admin > Data Streams > choose your stream > Measurement ID
		measurementID   string
		protocolVersion string
		httpClient      *http.Client
		middlewares     []MPv2Middleware
	}
	// MPv2Option configures an [MPv2] client.
	MPv2Option func(*MPv2)
	// MPv2Handler sends payload in the context of the incoming request r.
	MPv2Handler func(r *http.Request, payload *mpv2.Payload[any]) error
	// MPv2Middleware wraps an [MPv2Handler], typically to enrich the payload.
	MPv2Middleware func(next MPv2Handler) MPv2Handler
)

// ------------------------------------------------------------------------------------------------
// ~ Options
// ------------------------------------------------------------------------------------------------

// MPv2WithHTTPClient sets the HTTP client used for upstream requests.
// The default is [http.DefaultClient].
func MPv2WithHTTPClient(v *http.Client) MPv2Option {
	return func(o *MPv2) {
		o.httpClient = v
	}
}

// MPv2WithPath sets the collect path appended to the host.
// The default is "/mp/collect".
func MPv2WithPath(v string) MPv2Option {
	return func(o *MPv2) {
		o.path = v
	}
}

// MPv2WithCookies appends names of cookies forwarded from the incoming request.
// The defaults are gtm_auth, gtm_debug and gtm_preview.
func MPv2WithCookies(v ...string) MPv2Option {
	return func(o *MPv2) {
		o.cookies = append(o.cookies, v...)
	}
}

// MPv2WithAPISecret sets the api_secret query parameter.
// Create a secret in the Google Analytics UI under Admin > Data Streams >
// choose your stream > Measurement Protocol > Create.
func MPv2WithAPISecret(v string) MPv2Option {
	return func(o *MPv2) {
		o.apiSecret = v
	}
}

// MPv2WithMeasurementID sets the measurement_id query parameter, the identifier
// of a data stream, e.g. G-F9XM71K45T.
func MPv2WithMeasurementID(v string) MPv2Option {
	return func(o *MPv2) {
		o.measurementID = v
	}
}

// MPv2WithMiddlewares appends middlewares.
// Middlewares wrap in list order, so the last one runs first (outermost).
func MPv2WithMiddlewares(v ...MPv2Middleware) MPv2Option {
	return func(o *MPv2) {
		o.middlewares = append(o.middlewares, v...)
	}
}

// ------------------------------------------------------------------------------------------------
// ~ Constructor
// ------------------------------------------------------------------------------------------------

// NewMPv2 returns an [MPv2] client that sends to host. Custom middlewares from
// [MPv2WithMiddlewares] run inside the built-in client ID middleware.
func NewMPv2(l *zap.Logger, host string, opts ...MPv2Option) *MPv2 {
	inst := &MPv2{
		l:               l,
		host:            host,
		path:            "/mp/collect",
		cookies:         []string{"gtm_auth", "gtm_debug", "gtm_preview"},
		protocolVersion: "2",
		httpClient:      http.DefaultClient,
	}
	for _, opt := range opts {
		opt(inst)
	}

	inst.middlewares = append(inst.middlewares,
		MPv2MiddlewarClientID,
	)

	return inst
}

// ------------------------------------------------------------------------------------------------
// ~ Getter
// ------------------------------------------------------------------------------------------------

// HTTPClient returns the HTTP client used for upstream requests.
func (c *MPv2) HTTPClient() *http.Client {
	return c.httpClient
}

// ------------------------------------------------------------------------------------------------
// ~ Public methods
// ------------------------------------------------------------------------------------------------

// Collect sends events in one payload with [MPv2.SendPayload].
func (c *MPv2) Collect(r *http.Request, events ...sesamy.AnyEvent) error {
	payload := mpv2.NewPayload[any]()
	for _, event := range events {
		payload.Events = append(payload.Events, event.AnyEvent())
	}

	return c.SendPayload(r, payload)
}

// SendPayload runs payload through the middlewares and sends it with
// [MPv2.SendRaw]. r is the incoming request whose context, headers and cookies
// are used.
func (c *MPv2) SendPayload(r *http.Request, payload *mpv2.Payload[any]) error {
	next := c.SendRaw
	for _, middleware := range c.middlewares {
		next = middleware(next)
	}

	return next(r, payload)
}

// SendRaw posts payload as JSON to the tagging endpoint without running
// middlewares. It copies the headers of r except Cookie and Authorization and
// forwards only the configured cookies. It returns an error if the upstream
// response status is not 200 OK.
func (c *MPv2) SendRaw(r *http.Request, payload *mpv2.Payload[any]) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return errors.Wrap(err, "failed to encode payload")
	}

	req, err := http.NewRequestWithContext(
		r.Context(),
		http.MethodPost,
		fmt.Sprintf("%s%s", c.host, c.path),
		bytes.NewReader(jsonPayload),
	)
	if err != nil {
		return errors.Wrap(err, "failed to create request")
	}

	// query
	qry := req.URL.Query()
	if len(c.apiSecret) > 0 {
		qry.Add("api_secret", c.apiSecret)
	}

	if len(c.measurementID) > 0 {
		qry.Add("measurement_id", c.measurementID)
	}

	req.URL.RawQuery = qry.Encode()

	// copy headers without credentials, cookies are forwarded explicitly below
	req.Header = r.Header.Clone()
	req.Header.Del("Cookie")
	req.Header.Del("Authorization")
	req.Header.Set("Content-Type", "application/json")

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
