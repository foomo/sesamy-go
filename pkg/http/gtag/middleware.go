package gtag

import (
	"net/http"

	"github.com/foomo/sesamy-go/pkg/encoding/gtag"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type (
	Middleware        func(next MiddlewareHandler) MiddlewareHandler
	MiddlewareHandler func(l *zap.Logger, w http.ResponseWriter, r *http.Request, payload *gtag.Payload) error
)

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

func MiddlewareLogger(next MiddlewareHandler) MiddlewareHandler {
	return func(l *zap.Logger, w http.ResponseWriter, r *http.Request, payload *gtag.Payload) error {
		if spanCtx := trace.SpanContextFromContext(r.Context()); spanCtx.IsValid() && spanCtx.IsSampled() {
			l = l.With(zap.String("trace_id", spanCtx.TraceID().String()), zap.String("span_id", spanCtx.SpanID().String()))
		}
		l = l.With(
			zap.String("event_name", gtag.GetDefault(payload.EventName, "-").String()),
			zap.String("event_user_id", gtag.GetDefault(payload.UserID, "-")),
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
