package mpv2

import (
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/davecgh/go-spew/spew"
	"github.com/foomo/sesamy-go/pkg/encoding/mpv2"
	"go.uber.org/zap"
)

func PublisherMiddlewareDebugMode(next PublisherHandler) PublisherHandler {
	return func(l *zap.Logger, msg *message.Message) error {
		var payload *mpv2.Payload[any]
		if err := json.Unmarshal(msg.Payload, &payload); err != nil {
			return err
		}
		if payload.DebugMode {
			for i, event := range payload.Events {
				if params, ok := event.Params.(map[string]any); ok {
					params["debug_mode"] = "1"
				}
				payload.Events[i] = event
			}
			spew.Dump(payload.Events)
			out, err := json.Marshal(payload)
			if err != nil {
				return err
			}
			msg.Payload = out
		}
		return next(l, msg)
	}
}
