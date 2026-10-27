package collect

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/foomo/sesamy-go/pkg/encoding/gtag"
	"github.com/foomo/sesamy-go/pkg/encoding/mpv2"
	gtaghttp "github.com/foomo/sesamy-go/pkg/http/gtag"
	mpv2http "github.com/foomo/sesamy-go/pkg/http/mpv2"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type (
	// Collect receives gtag.js and Measurement Protocol v2 requests and forwards
	// them to a server-side tagging endpoint. Create one with [New].
	Collect struct {
		l                   *zap.Logger
		taggingURL          string
		taggingClient       *http.Client
		measurementID       string
		gtagHTTPMiddlewares []gtaghttp.Middleware
		mpv2HTTPMiddlewares []mpv2http.Middleware
	}
	// Option configures a [Collect].
	Option func(*Collect) error
)

// ------------------------------------------------------------------------------------------------
// ~ Options
// ------------------------------------------------------------------------------------------------

// WithTagging sets the base URL of the server-side tagging endpoint.
// The request path is appended to it.
func WithTagging(v string) Option {
	return func(c *Collect) error {
		c.taggingURL = v
		return nil
	}
}

// WithTaggingClient sets the HTTP client used for upstream requests.
// The default is [http.DefaultClient].
func WithTaggingClient(v *http.Client) Option {
	return func(c *Collect) error {
		c.taggingClient = v
		return nil
	}
}

// WithMeasurementID sets the default `measurement_id` for MPv2 requests that do not provide one.
func WithMeasurementID(v string) Option {
	return func(c *Collect) error {
		c.measurementID = v
		return nil
	}
}

// WithGTagHTTPMiddlewares appends middlewares to the gtag handler.
// Middlewares wrap in list order, so the last one runs first (outermost).
// Put e.g. MiddlewareLogger before the middlewares that fill the payload.
func WithGTagHTTPMiddlewares(v ...gtaghttp.Middleware) Option {
	return func(c *Collect) error {
		c.gtagHTTPMiddlewares = append(c.gtagHTTPMiddlewares, v...)
		return nil
	}
}

// WithMPv2HTTPMiddlewares appends middlewares to the MPv2 handler.
// Middlewares wrap in list order, so the last one runs first (outermost).
// Put e.g. MiddlewareLogger before the middlewares that fill the payload.
func WithMPv2HTTPMiddlewares(v ...mpv2http.Middleware) Option {
	return func(c *Collect) error {
		c.mpv2HTTPMiddlewares = append(c.mpv2HTTPMiddlewares, v...)
		return nil
	}
}

// ------------------------------------------------------------------------------------------------
// ~ Constructor
// ------------------------------------------------------------------------------------------------

// New returns a [Collect] configured by opts. Nil options are skipped.
// It returns the first error an option returns.
func New(l *zap.Logger, opts ...Option) (*Collect, error) {
	inst := &Collect{
		l:             l,
		taggingClient: http.DefaultClient,
	}

	for _, opt := range opts {
		if opt != nil {
			if err := opt(inst); err != nil {
				return nil, err
			}
		}
	}

	return inst, nil
}

// ------------------------------------------------------------------------------------------------
// ~ Public methods
// ------------------------------------------------------------------------------------------------

// GTagHTTPHandler decodes a gtag.js request into one payload per event (see
// [gtaghttp.HandlerBatch]), runs each through the gtag middlewares and forwards
// it upstream. Only the first upstream response is written to w.
func (c *Collect) GTagHTTPHandler(w http.ResponseWriter, r *http.Request) {
	// retrieve payloads
	payloads := gtaghttp.HandlerBatch(w, r)
	if payloads == nil {
		http.Error(w, "failed to decode payload", http.StatusBadRequest)
		return
	}

	// compose middlewares
	next := c.gtagHandler
	for _, middleware := range c.gtagHTTPMiddlewares {
		next = middleware(next)
	}

	// run handler per event, only the first upstream response reaches the client
	for i, payload := range payloads {
		var rw = w
		if i > 0 {
			rw = &discardResponseWriter{header: http.Header{}}
		}

		if err := next(c.l, rw, r, payload); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// MPv2HTTPHandler decodes a Measurement Protocol v2 request, runs it through the
// MPv2 middlewares and forwards it upstream, writing the upstream response to w.
func (c *Collect) MPv2HTTPHandler(w http.ResponseWriter, r *http.Request) {
	// retrieve payload
	payload := mpv2http.Handler(w, r)

	// compose middlewares
	next := c.mpv2Handler
	for _, middleware := range c.mpv2HTTPMiddlewares {
		next = middleware(next)
	}

	// run handler
	if err := next(c.l, w, r, payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ------------------------------------------------------------------------------------------------
// ~ Private methods
// ------------------------------------------------------------------------------------------------

func (c *Collect) gtagHandler(l *zap.Logger, w http.ResponseWriter, r *http.Request, payload *gtag.Payload) error {
	query, body, err := gtag.Encode(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(r.Context(), http.MethodPost, fmt.Sprintf("%s%s?%s", c.taggingURL, "/g/collect", query), body) //nolint:gosec // G704
	if err != nil {
		return errors.Wrap(err, "failed to create request")
	}

	// copy headers
	req.Header = r.Header.Clone()

	resp, err := c.taggingClient.Do(req) //nolint:gosec // taggingURL is a trusted config value set at construction time
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// copy headers
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(resp.StatusCode)

	if _, err := io.Copy(w, resp.Body); err != nil {
		return err
	}

	return nil
}

func (c *Collect) mpv2Handler(l *zap.Logger, w http.ResponseWriter, r *http.Request, payload *mpv2.Payload[any]) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(r.Context(), http.MethodPost, fmt.Sprintf("%s%s", c.taggingURL, "/mp/collect"), bytes.NewReader(body))
	if err != nil {
		return errors.Wrap(err, "failed to create request")
	}

	// copy headers
	req.Header = r.Header.Clone()
	// copy raw query
	req.URL.RawQuery = r.URL.RawQuery
	if query := r.URL.Query(); c.measurementID != "" && query.Get("measurement_id") == "" {
		query.Set("measurement_id", c.measurementID)
		req.URL.RawQuery = query.Encode()
	}

	resp, err := c.taggingClient.Do(req) //nolint:gosec // taggingURL is a trusted config value set at construction time
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// copy headers
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(resp.StatusCode)

	if _, err := io.Copy(w, resp.Body); err != nil {
		return err
	}

	return nil
}

// ------------------------------------------------------------------------------------------------
// ~ Private types
// ------------------------------------------------------------------------------------------------

// discardResponseWriter is an [http.ResponseWriter] that drops everything
// written to it.
type discardResponseWriter struct {
	header http.Header
}

func (w *discardResponseWriter) Header() http.Header         { return w.header }
func (w *discardResponseWriter) Write(b []byte) (int, error) { return len(b), nil }
func (w *discardResponseWriter) WriteHeader(int)             {}
