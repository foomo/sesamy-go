package mpv2

import (
	"net/http"
	"strings"
	"time"

	"github.com/foomo/sesamy-go/pkg/encoding/mpv2"
	"github.com/foomo/sesamy-go/pkg/session"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

func SubscriberMiddlewareSessionID(measurementID string) SubscriberMiddleware {
	measurementID = strings.Split(measurementID, "-")[1]
	return func(next SubscriberHandler) SubscriberHandler {
		return func(l *zap.Logger, r *http.Request, payload *mpv2.Payload[any]) error {
			if payload.SessionID == "" {
				value, err := session.ParseGASessionID(r, measurementID)
				if err != nil && !errors.Is(err, http.ErrNoCookie) {
					return errors.Wrap(err, "failed to parse client cookie")
				}
				payload.SessionID = value
			}
			return next(l, r, payload)
		}
	}
}

func SubscriberMiddlewareClientID(next SubscriberHandler) SubscriberHandler {
	return func(l *zap.Logger, r *http.Request, payload *mpv2.Payload[any]) error {
		if payload.ClientID == "" {
			value, err := session.ParseGAClientID(r)
			if err != nil {
				return err
			}
			payload.ClientID = value
		}
		return next(l, r, payload)
	}
}

func SubscriberMiddlewareDebugMode(next SubscriberHandler) SubscriberHandler {
	return func(l *zap.Logger, r *http.Request, payload *mpv2.Payload[any]) error {
		if !payload.DebugMode {
			payload.DebugMode = session.IsGTMDebug(r)
		}
		return next(l, r, payload)
	}
}

func SubscriberMiddlewareUserID(cookieName string) SubscriberMiddleware {
	return func(next SubscriberHandler) SubscriberHandler {
		return func(l *zap.Logger, r *http.Request, payload *mpv2.Payload[any]) error {
			if payload.UserID == "" {
				value, err := r.Cookie(cookieName)
				if err != nil && !errors.Is(err, http.ErrNoCookie) {
					return err
				}
				payload.UserID = value.Value
			}
			return next(l, r, payload)
		}
	}
}

func SubscriberMiddlewareTimestamp(next SubscriberHandler) SubscriberHandler {
	return func(l *zap.Logger, r *http.Request, payload *mpv2.Payload[any]) error {
		if payload.TimestampMicros == 0 {
			payload.TimestampMicros = time.Now().UnixMicro()
		}
		return next(l, r, payload)
	}
}

func SubscriberMiddlewareLogger(next SubscriberHandler) SubscriberHandler {
	return func(l *zap.Logger, r *http.Request, payload *mpv2.Payload[any]) error {
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
		)

		err := next(l, r, payload)
		if err != nil {
			l.Error("handled event", zap.Error(err))
		} else {
			l.Info("handled event")
		}
		return err
	}
}
