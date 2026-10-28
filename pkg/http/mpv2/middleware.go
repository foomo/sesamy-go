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
	// MiddlewareHandler handles a decoded Measurement Protocol v2 payload of request r.
	MiddlewareHandler func(l *zap.Logger, w http.ResponseWriter, r *http.Request, payload *mpv2.Payload[any]) error
	// Middleware wraps a [MiddlewareHandler].
	Middleware func(next MiddlewareHandler) MiddlewareHandler
)

// MiddlewareEventHandler calls h for each event of the payload.
// An error from h aborts the chain.
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

// MiddlewareSessionID sets an empty session ID and the ga_session_id and
// ga_session_number event params from the _ga_<suffix> cookie, where suffix is
// measurementID without its prefix (see [session.MeasurementIDSuffix]).
// A missing cookie is ignored; an invalid cookie value returns an error.
func MiddlewareSessionID(measurementID string) Middleware {
	measurementID = session.MeasurementIDSuffix(measurementID)

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

// MiddlewareClientID sets an empty client ID from the _ga cookie.
// A missing cookie is ignored; an invalid cookie value returns an error.
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

// MiddlewareDebugMode enables debug mode if the request has a gtm_debug cookie.
func MiddlewareDebugMode(next MiddlewareHandler) MiddlewareHandler {
	return func(l *zap.Logger, w http.ResponseWriter, r *http.Request, payload *mpv2.Payload[any]) error {
		if !payload.DebugMode && session.IsGTMDebug(r) {
			payload.DebugMode = true
		}

		return next(l, w, r, payload)
	}
}

// MiddlewareUserID sets an empty user ID from the cookie cookieName.
// A missing cookie is ignored.
func MiddlewareUserID(cookieName string) Middleware {
	return func(next MiddlewareHandler) MiddlewareHandler {
		return func(l *zap.Logger, w http.ResponseWriter, r *http.Request, payload *mpv2.Payload[any]) error {
			if payload.UserID == "" {
				value, err := r.Cookie(cookieName)
				if err != nil && !errors.Is(err, http.ErrNoCookie) {
					return err
				}

				if value != nil {
					payload.UserID = value.Value
				}
			}

			return next(l, w, r, payload)
		}
	}
}

// MiddlewareTimestamp sets an empty timestamp to the current time.
func MiddlewareTimestamp(next MiddlewareHandler) MiddlewareHandler {
	return func(l *zap.Logger, w http.ResponseWriter, r *http.Request, payload *mpv2.Payload[any]) error {
		if payload.TimestampMicros == 0 {
			payload.TimestampMicros = time.Now().UnixMicro()
		}

		return next(l, w, r, payload)
	}
}

// MiddlewareUserAgent sets the user_agent event param from the User-Agent
// header unless already set.
func MiddlewareUserAgent(next MiddlewareHandler) MiddlewareHandler {
	return func(l *zap.Logger, w http.ResponseWriter, r *http.Request, payload *mpv2.Payload[any]) error {
		if userAgent := r.Header.Get("User-Agent"); userAgent != "" {
			for i, event := range payload.Events {
				if value, ok := event.Params.(map[string]any); ok {
					if value["user_agent"] == nil {
						value["user_agent"] = userAgent
					}

					event.Params = value
				}

				payload.Events[i] = event
			}
		}

		return next(l, w, r, payload)
	}
}

// MiddlewareIPOverride sets the ip_override event param unless already set.
// The value is the first non-empty header of CF-Connecting-IP,
// X-Original-Forwarded-For, X-Forwarded-For and X-Real-Ip.
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
					if value["ip_override"] == nil {
						value["ip_override"] = ipOverride
					}

					event.Params = value
				}

				payload.Events[i] = event
			}
		}

		return next(l, w, r, payload)
	}
}

// MiddlewareEngagementTime sets the engagement_time_msec event param to 100
// unless already set.
func MiddlewareEngagementTime(next MiddlewareHandler) MiddlewareHandler {
	return func(l *zap.Logger, w http.ResponseWriter, r *http.Request, payload *mpv2.Payload[any]) error {
		for i, event := range payload.Events {
			if value, ok := event.Params.(map[string]any); ok {
				if value["engagement_time_msec"] == nil {
					value["engagement_time_msec"] = 100
				}

				event.Params = value
			}

			payload.Events[i] = event
		}

		return next(l, w, r, payload)
	}
}

// MiddlewarePageLocation sets the page_title, page_referrer and page_location
// event params from the X-Page-Title, X-Page-Referrer and Referer headers
// unless already set.
func MiddlewarePageLocation(next MiddlewareHandler) MiddlewareHandler {
	return func(l *zap.Logger, w http.ResponseWriter, r *http.Request, payload *mpv2.Payload[any]) error {
		pageTitle := r.Header.Get("X-Page-Title")
		pageLocation := r.Header.Get("Referer")
		pageReferrer := r.Header.Get("X-Page-Referrer")

		for i, event := range payload.Events {
			if value, ok := event.Params.(map[string]any); ok {
				if value["page_title"] == nil && pageTitle != "" {
					value["page_title"] = pageTitle
				}

				if value["page_referrer"] == nil && pageReferrer != "" {
					value["page_referrer"] = pageReferrer
				}

				if value["page_location"] == nil && pageLocation != "" {
					value["page_location"] = pageLocation
				}

				event.Params = value
			}

			payload.Events[i] = event
		}

		return next(l, w, r, payload)
	}
}

// MiddlewareWithTimeout replaces the request context with one that ignores
// cancellation of the incoming request and expires after timeout.
func MiddlewareWithTimeout(timeout time.Duration) Middleware {
	return func(next MiddlewareHandler) MiddlewareHandler {
		return func(l *zap.Logger, w http.ResponseWriter, r *http.Request, payload *mpv2.Payload[any]) error {
			ctx, cancel := context.WithTimeout(context.WithoutCancel(r.Context()), timeout)
			defer cancel()

			return next(l, w, r.WithContext(ctx), payload)
		}
	}
}

// MiddlewareLogger adds event and trace fields to the logger passed down the
// chain and logs the result of next.
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
