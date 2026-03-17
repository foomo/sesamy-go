package gtag

import (
	"context"
	"net/http"
	"time"

	"github.com/foomo/sesamy-go/pkg/encoding/gtag"
	"github.com/foomo/sesamy-go/pkg/encoding/gtagencode"
	"github.com/foomo/sesamy-go/pkg/encoding/mpv2"
	"github.com/foomo/sesamy-go/pkg/encoding/mpv2encode"
	sesamyhttp "github.com/foomo/sesamy-go/pkg/http"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type (
	Middleware        func(next MiddlewareHandler) MiddlewareHandler
	MiddlewareHandler func(l *zap.Logger, w http.ResponseWriter, r *http.Request, payload *gtag.Payload) error
)

func MiddlewareEventHandler(h sesamyhttp.EventHandler) Middleware {
	return func(next MiddlewareHandler) MiddlewareHandler {
		return func(l *zap.Logger, w http.ResponseWriter, r *http.Request, payload *gtag.Payload) error {
			var mpv2Payload *mpv2.Payload[any]
			if err := gtagencode.MPv2(*payload, &mpv2Payload); err != nil {
				return errors.Wrap(err, "failed to encode gtag to mpv2")
			}

			for i, event := range mpv2Payload.Events {
				if err := h(l, r, &event); err != nil {
					return err
				}
				mpv2Payload.Events[i] = event
			}

			if err := mpv2encode.GTag[any](*mpv2Payload, &payload); err != nil {
				return errors.Wrap(err, "failed to encode mpv2 to gtag")
			}
			return next(l, w, r, payload)
		}
	}
}

func MiddlewareUserID(cookieName string) Middleware {
	return func(next MiddlewareHandler) MiddlewareHandler {
		return func(l *zap.Logger, w http.ResponseWriter, r *http.Request, payload *gtag.Payload) error {
			if cookie, err := r.Cookie(cookieName); err == nil {
				payload.UserID = gtag.Set(cookie.Value)
			}
			return next(l, w, r, payload)
		}
	}
}

func MiddlewareWithTimeout(timeout time.Duration) Middleware {
	return func(next MiddlewareHandler) MiddlewareHandler {
		return func(l *zap.Logger, w http.ResponseWriter, r *http.Request, payload *gtag.Payload) error {
			ctx, cancel := context.WithTimeout(context.WithoutCancel(r.Context()), timeout)
			defer cancel()
			return next(l, w, r.WithContext(ctx), payload)
		}
	}
}

func MiddlewareLogger(next MiddlewareHandler) MiddlewareHandler {
	return func(l *zap.Logger, w http.ResponseWriter, r *http.Request, payload *gtag.Payload) error {
		if spanCtx := trace.SpanContextFromContext(r.Context()); spanCtx.IsValid() && spanCtx.IsSampled() {
			l = l.With(zap.String("trace_id", spanCtx.TraceID().String()), zap.String("span_id", spanCtx.SpanID().String()))
		}
		l = l.With(
			zap.String("event_name", gtag.GetDefault(payload.EventName, "-").String()),
			zap.String("event_user_id", gtag.GetDefault(payload.UserID, "-")),
			zap.String("event_client_id", gtag.GetDefault(payload.ClientID, "-")),
			zap.String("event_session_id", gtag.GetDefault(payload.SessionID, "-")),
		)
		err := next(l, w, r, payload)
		if err != nil {
			l.Error("handled event", zap.Error(err))
		} else {
			l.Info("handled event")
		}
		return err
	}
}
