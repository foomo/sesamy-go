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
	Collect struct {
		l                   *zap.Logger
		taggingURL          string
		taggingClient       *http.Client
		gtagHTTPMiddlewares []gtaghttp.Middleware
		mpv2HTTPMiddlewares []mpv2http.Middleware
	}
	Option func(*Collect) error
)

// ------------------------------------------------------------------------------------------------
// ~ Options
// ------------------------------------------------------------------------------------------------

func WithTagging(v string) Option {
	return func(c *Collect) error {
		c.taggingURL = v
		return nil
	}
}

func WithTaggingClient(v *http.Client) Option {
	return func(c *Collect) error {
		c.taggingClient = v
		return nil
	}
}

func WithGTagHTTPMiddlewares(v ...gtaghttp.Middleware) Option {
	return func(c *Collect) error {
		c.gtagHTTPMiddlewares = append(c.gtagHTTPMiddlewares, v...)
		return nil
	}
}

func WithMPv2HTTPMiddlewares(v ...mpv2http.Middleware) Option {
	return func(c *Collect) error {
		c.mpv2HTTPMiddlewares = append(c.mpv2HTTPMiddlewares, v...)
		return nil
	}
}

// ------------------------------------------------------------------------------------------------
// ~ Constructor
// ------------------------------------------------------------------------------------------------

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

func (c *Collect) GTagHTTPHandler(w http.ResponseWriter, r *http.Request) {
	// retrieve payload
	payload := gtaghttp.Handler(w, r)

	// compose middlewares
	next := c.gtagHandler
	for _, middleware := range c.gtagHTTPMiddlewares {
		next = middleware(next)
	}

	// run handler
	if err := next(c.l, w, r, payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

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
	values, body, err := gtag.Encode(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(r.Context(), http.MethodPost, fmt.Sprintf("%s%s?%s", c.taggingURL, "/g/collect", gtag.EncodeValues(values)), body)
	if err != nil {
		return errors.Wrap(err, "failed to create request")
	}

	// copy headers
	req.Header = r.Header.Clone()

	resp, err := c.taggingClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// copy headers
	r.Header = resp.Header.Clone()

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

	resp, err := c.taggingClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// copy headers
	r.Header = resp.Header.Clone()

	if _, err := io.Copy(w, resp.Body); err != nil {
		return err
	}

	return nil
}
