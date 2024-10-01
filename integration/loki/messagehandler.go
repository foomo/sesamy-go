package loki

import (
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/foomo/sesamy-go/pkg/encoding/mpv2"
	"github.com/pkg/errors"
)

func MPv2MessageHandler(loki *Loki) message.NoPublishHandlerFunc {
	return func(msg *message.Message) error {
		var payload mpv2.Payload[any]

		// unmarshal payload
		if err := json.Unmarshal(msg.Payload, &payload); err != nil {
			return errors.Wrap(err, "failed to unmarshal payload")
		}

		loki.Write(payload)
		return nil
	}
}
