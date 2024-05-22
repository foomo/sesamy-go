package mpv2

import (
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/foomo/sesamy-go/pkg/encoding/mpv2"
	"github.com/pkg/errors"
)

func MessageHandler(handler func(payload *mpv2.Payload[any], msg *message.Message) error) func(msg *message.Message) ([]*message.Message, error) {
	return func(msg *message.Message) ([]*message.Message, error) {
		var payload *mpv2.Payload[any]

		// unmarshal payload
		if err := json.Unmarshal(msg.Payload, &payload); err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal payload")
		}

		// handle payload
		if err := handler(payload, msg); err != nil {
			return nil, err
		}

		// marshal payload
		b, err := json.Marshal(payload)
		if err != nil {
			return nil, errors.Wrap(err, "failed to marshal payload")
		}
		msg.Payload = b

		return []*message.Message{msg}, nil
	}
}
