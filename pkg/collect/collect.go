package collect

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/foomo/sesamy-go/pkg/encoding/gtag"
	"github.com/foomo/sesamy-go/pkg/encoding/gtagencode"
	"github.com/foomo/sesamy-go/pkg/encoding/mpv2"
	"github.com/foomo/sesamy-go/pkg/encoding/mpv2encode"
	sesamyhttp "github.com/foomo/sesamy-go/pkg/http"
	gtaghttp "github.com/foomo/sesamy-go/pkg/http/gtag"
	mpv2http "github.com/foomo/sesamy-go/pkg/http/mpv2"
	"github.com/foomo/sesamy-go/pkg/sesamy"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type (
	Collect struct {
		l               *zap.Logger
		gtagProxy       *httputil.ReverseProxy
		mpv2Proxy       *httputil.ReverseProxy
		gtagMiddlewares []gtaghttp.Middleware
		mpv2Middlewares []mpv2http.Middleware
		eventHandlers   []sesamyhttp.EventHandler
	}
	Option func(*Collect) error
)

// ------------------------------------------------------------------------------------------------
// ~ Options
// ------------------------------------------------------------------------------------------------

func WithGTag(endpoint string) Option {
	return func(c *Collect) error {
		target, err := url.Parse(endpoint)
		if err != nil {
			return err
		}
		proxy := httputil.NewSingleHostReverseProxy(target)
		c.gtagProxy = proxy
		return nil
	}
}

func WithMPv2(endpoint string) Option {
	return func(c *Collect) error {
		target, err := url.Parse(endpoint)
		if err != nil {
			return err
		}
		proxy := httputil.NewSingleHostReverseProxy(target)
		c.mpv2Proxy = proxy
		return nil
	}
}

func WithGTagMiddlewares(v ...gtaghttp.Middleware) Option {
	return func(c *Collect) error {
		c.gtagMiddlewares = append(c.gtagMiddlewares, v...)
		return nil
	}
}

func WithMPv2Middlewares(v ...mpv2http.Middleware) Option {
	return func(c *Collect) error {
		c.mpv2Middlewares = append(c.mpv2Middlewares, v...)
		return nil
	}
}

func WithEventHandlers(v ...sesamyhttp.EventHandler) Option {
	return func(c *Collect) error {
		c.eventHandlers = append(c.eventHandlers, v...)
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
	for _, middleware := range c.gtagMiddlewares {
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
	for _, middleware := range c.mpv2Middlewares {
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
	var mpv2Payload *mpv2.Payload[any]
	if err := gtagencode.MPv2(*payload, &mpv2Payload); err != nil {
		return errors.Wrap(err, "failed to encode gtag to mpv2")
	}

	for i, event := range mpv2Payload.Events {
		if err := c.mpv2EventHandler(r, &event); err != nil {
			return err
		}
		mpv2Payload.Events[i] = event
	}

	if err := mpv2encode.GTag[any](*mpv2Payload, &payload); err != nil {
		return errors.Wrap(err, "failed to encode mpv2 to gtag")
	}

	if c.gtagProxy == nil {
		c.gtagProxy.ServeHTTP(w, r)
	}
	return nil
}

func (c *Collect) mpv2Handler(l *zap.Logger, w http.ResponseWriter, r *http.Request, payload *mpv2.Payload[any]) error {
	for i, event := range payload.Events {
		if err := c.mpv2EventHandler(r, &event); err != nil {
			return err
		}
		payload.Events[i] = event
	}

	if c.mpv2Proxy == nil {
		c.mpv2Proxy.ServeHTTP(w, r)
	}
	return nil
}

func (c *Collect) mpv2EventHandler(r *http.Request, event *sesamy.Event[any]) error {
	for _, handler := range c.eventHandlers {
		if err := handler(r, event); err != nil {
			return err
		}
	}
	return nil
}
