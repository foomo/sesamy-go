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
			spew.Dump(payload.Events)
		}
		return next(l, msg)
	}
}
