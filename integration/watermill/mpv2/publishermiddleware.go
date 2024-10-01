package mpv2

import (
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/foomo/sesamy-go/pkg/encoding/mpv2"
	"go.uber.org/zap"
)

func PublisherMiddlewareIgnoreError(next PublisherHandler) PublisherHandler {
	return func(l *zap.Logger, msg *message.Message) error {
		err := next(l, msg)
		l.With(zap.Error(err)).Warn("ignoring error")
		return nil
	}
}

func PublisherMiddlewareEventParams(next PublisherHandler) PublisherHandler {
	return func(l *zap.Logger, msg *message.Message) error {
		var payload *mpv2.Payload[any]
		if err := json.Unmarshal(msg.Payload, &payload); err != nil {
			return err
		}
		for i, event := range payload.Events {
			if params, ok := event.Params.(map[string]any); ok {
				if payload.DebugMode {
					params["debug_mode"] = "1"
					payload.DebugMode = false
				}
				if len(payload.SessionID) > 0 {
					params["session_id"] = payload.SessionID
					payload.SessionID = ""
				}
				if payload.EngagementTimeMSec > 0 {
					params["engagement_time_msec"] = payload.EngagementTimeMSec
					payload.EngagementTimeMSec = 0
				}
				event.Params = params
			}
			payload.Events[i] = event

			out, err := json.Marshal(payload)
			if err != nil {
				return err
			}

			msg.Payload = out
		}
		return next(l, msg)
	}
}
