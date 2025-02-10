package mpv2

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/foomo/sesamy-go/pkg/encoding/mpv2"
	sesamyhttp "github.com/foomo/sesamy-go/pkg/http"
	"github.com/foomo/sesamy-go/pkg/session"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type (
	MiddlewareHandler func(l *zap.Logger, w http.ResponseWriter, r *http.Request, payload *mpv2.Payload[any]) error
	Middleware        func(next MiddlewareHandler) MiddlewareHandler
)

func MiddlewareEventHandler(h sesamyhttp.EventHandler) Middleware {
	return func(next MiddlewareHandler) MiddlewareHandler {
		return func(l *zap.Logger, w http.ResponseWriter, r *http.Request, payload *mpv2.Payload[any]) error {
			for i, event := range payload.Events {
				if err := h(l, r, &event); err != nil {
					return err
				}
				payload.Events[i] = event
			}
			return next(l, w, r, payload)
		}
	}
}

func MiddlewareSessionID(measurementID string) Middleware {
	measurementID = strings.Split(measurementID, "-")[1]
	return func(next MiddlewareHandler) MiddlewareHandler {
		return func(l *zap.Logger, w http.ResponseWriter, r *http.Request, payload *mpv2.Payload[any]) error {
			if payload.SessionID == "" {
				id, err := session.ParseGASessionID(r, measurementID)
				if err != nil && !errors.Is(err, http.ErrNoCookie) {
					return err
				}

				number, err := session.ParseGASessionNumber(r, measurementID)
				if err != nil && !errors.Is(err, http.ErrNoCookie) {
					return err
				}

				payload.SessionID = id
				for i, event := range payload.Events {
					if value, ok := event.Params.(map[string]any); ok {
						value["ga_session_id"] = id
						value["ga_session_number"] = number
						event.Params = value
					}
					payload.Events[i] = event
				}
			}
			return next(l, w, r, payload)
		}
	}
}

func MiddlewareClientID(next MiddlewareHandler) MiddlewareHandler {
	return func(l *zap.Logger, w http.ResponseWriter, r *http.Request, payload *mpv2.Payload[any]) error {
		if payload.ClientID == "" {
			value, err := session.ParseGAClientID(r)
			if err != nil && !errors.Is(err, http.ErrNoCookie) {
				return err
			}
			payload.ClientID = value
		}
		return next(l, w, r, payload)
	}
}

func MiddlewareDebugMode(next MiddlewareHandler) MiddlewareHandler {
	return func(l *zap.Logger, w http.ResponseWriter, r *http.Request, payload *mpv2.Payload[any]) error {
		if !payload.DebugMode && session.IsGTMDebug(r) {
			payload.DebugMode = true
			for i, event := range payload.Events {
				if value, ok := event.Params.(map[string]any); ok {
					value["debug_mode"] = true
					event.Params = value
				}
				payload.Events[i] = event
			}
		}
		return next(l, w, r, payload)
	}
}

func MiddlewareUserID(cookieName string) Middleware {
	return func(next MiddlewareHandler) MiddlewareHandler {
		return func(l *zap.Logger, w http.ResponseWriter, r *http.Request, payload *mpv2.Payload[any]) error {
			if payload.UserID == "" {
				value, err := r.Cookie(cookieName)
				if err != nil && !errors.Is(err, http.ErrNoCookie) {
					return err
				}
				payload.UserID = value.Value
			}
			return next(l, w, r, payload)
		}
	}
}

func MiddlewareTimestamp(next MiddlewareHandler) MiddlewareHandler {
	return func(l *zap.Logger, w http.ResponseWriter, r *http.Request, payload *mpv2.Payload[any]) error {
		if payload.TimestampMicros == 0 {
			payload.TimestampMicros = time.Now().UnixMicro()
		}
		return next(l, w, r, payload)
	}
}

func MiddlewareUserAgent(next MiddlewareHandler) MiddlewareHandler {
	return func(l *zap.Logger, w http.ResponseWriter, r *http.Request, payload *mpv2.Payload[any]) error {
		if userAgent := r.Header.Get("User-Agent"); userAgent != "" {
			for i, event := range payload.Events {
				if value, ok := event.Params.(map[string]any); ok {
					value["user_agent"] = userAgent
					payload.Events[i] = event
				}
			}
		}
		return next(l, w, r, payload)
	}
}

func MiddlewareIPOverride(next MiddlewareHandler) MiddlewareHandler {
	return func(l *zap.Logger, w http.ResponseWriter, r *http.Request, payload *mpv2.Payload[any]) error {
		var ipOverride string
		for _, key := range []string{"CF-Connecting-IP", "X-Original-Forwarded-For", "X-Forwarded-For", "X-Real-Ip"} {
			if value := r.Header.Get(key); value != "" {
				ipOverride = value
				break
			}
		}
		if ipOverride != "" {
			for i, event := range payload.Events {
				if value, ok := event.Params.(map[string]any); ok {
					value["ip_override"] = ipOverride
					payload.Events[i] = event
				}
			}
		}
		return next(l, w, r, payload)
	}
}

func MiddlewarePageLocation(next MiddlewareHandler) MiddlewareHandler {
	return func(l *zap.Logger, w http.ResponseWriter, r *http.Request, payload *mpv2.Payload[any]) error {
		if referrer := r.Header.Get("Referer"); referrer != "" {
			for i, event := range payload.Events {
				if value, ok := event.Params.(map[string]any); ok {
					value["page_location"] = referrer
					payload.Events[i] = event
				}
			}
		}
		return next(l, w, r, payload)
	}
}

func MiddlewareWithTimeout(timeout time.Duration) Middleware {
	return func(next MiddlewareHandler) MiddlewareHandler {
		return func(l *zap.Logger, w http.ResponseWriter, r *http.Request, payload *mpv2.Payload[any]) error {
			ctx, cancel := context.WithTimeout(context.WithoutCancel(r.Context()), timeout)
			defer cancel()
			return next(l, w, r.WithContext(ctx), payload)
		}
	}
}

func MiddlewareLogger(next MiddlewareHandler) MiddlewareHandler {
	return func(l *zap.Logger, w http.ResponseWriter, r *http.Request, payload *mpv2.Payload[any]) error {
		eventNames := make([]string, len(payload.Events))
		for i, event := range payload.Events {
			eventNames[i] = event.Name.String()
		}

		if spanCtx := trace.SpanContextFromContext(r.Context()); spanCtx.IsValid() && spanCtx.IsSampled() {
			l = l.With(zap.String("trace_id", spanCtx.TraceID().String()), zap.String("span_id", spanCtx.SpanID().String()))
		}
		l = l.With(
			zap.String("event_names", strings.Join(eventNames, ",")),
			zap.String("event_user_id", payload.UserID),
			zap.String("event_client_id", payload.ClientID),
			zap.String("event_session_id", payload.SessionID),
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
