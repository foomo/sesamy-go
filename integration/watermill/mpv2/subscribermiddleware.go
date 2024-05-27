package mpv2

import (
	"net/http"
	"strings"
	"time"

	"github.com/foomo/sesamy-go/pkg/encoding/mpv2"
	"github.com/foomo/sesamy-go/pkg/session"
	"go.uber.org/zap"
)

func SubscriberMiddlewareClientID(next SubscriberHandler) SubscriberHandler {
	return func(l *zap.Logger, r *http.Request, payload *mpv2.Payload[any]) error {
		if payload.ClientID == "" {
			clientID, err := session.ParseGAClientID(r)
			if err != nil {
				return err
			}
			payload.ClientID = clientID
		}
		return next(l, r, payload)
	}
}

func SubscriberMiddlewareUserID(cookieName string) SubscriberMiddleware {
	return func(next SubscriberHandler) SubscriberHandler {
		return func(l *zap.Logger, r *http.Request, payload *mpv2.Payload[any]) error {
			if cookie, err := r.Cookie(cookieName); err == nil {
				payload.UserID = cookie.Value
			}
			return next(l, r, payload)
		}
	}
}

func SubscriberMiddlewareDebugMode(next SubscriberHandler) SubscriberHandler {
	return func(l *zap.Logger, r *http.Request, payload *mpv2.Payload[any]) error {
		if session.IsGTMDebug(r) {
			payload.DebugMode = true
		}
		return next(l, r, payload)
	}
}

func SubscriberMiddlewareTimestamp(next SubscriberHandler) SubscriberHandler {
	return func(l *zap.Logger, r *http.Request, payload *mpv2.Payload[any]) error {
		payload.TimestampMicros = time.Now().UnixMicro()
		return next(l, r, payload)
	}
}

func SubscriberMiddlewareLogger(next SubscriberHandler) SubscriberHandler {
	return func(l *zap.Logger, r *http.Request, payload *mpv2.Payload[any]) error {
		eventNames := make([]string, len(payload.Events))
		for i, event := range payload.Events {
			eventNames[i] = event.Name.String()
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
