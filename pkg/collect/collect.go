package collect

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/foomo/sesamy-go/pkg/encoding/gtag"
	"github.com/foomo/sesamy-go/pkg/encoding/mpv2"
	gtaghttp "github.com/foomo/sesamy-go/pkg/http/gtag"
	mpv2http "github.com/foomo/sesamy-go/pkg/http/mpv2"
	"go.uber.org/zap"
)

type (
	Collect struct {
		l                   *zap.Logger
		taggingProxy        *httputil.ReverseProxy
		gtagHTTPMiddlewares []gtaghttp.Middleware
		mpv2HTTPMiddlewares []mpv2http.Middleware
	}
	Option func(*Collect) error
)

// ------------------------------------------------------------------------------------------------
// ~ Options
// ------------------------------------------------------------------------------------------------

func WithTagging(endpoint string) Option {
	return func(c *Collect) error {
		target, err := url.Parse(endpoint)
		if err != nil {
			return err
		}
		c.l.Info("--->" + endpoint)
		proxy := httputil.NewSingleHostReverseProxy(target)
		proxy.ErrorLog = zap.NewStdLog(c.l)
		c.taggingProxy = proxy
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
		l: l,
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
	if c.taggingProxy != nil {
		c.taggingProxy.ServeHTTP(w, r)
	}
	return nil
}

func (c *Collect) mpv2Handler(l *zap.Logger, w http.ResponseWriter, r *http.Request, payload *mpv2.Payload[any]) error {
	if c.taggingProxy != nil {
		c.taggingProxy.ServeHTTP(w, r)
	}
	return nil
}
