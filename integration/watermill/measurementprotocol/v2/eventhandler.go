package v2

import (
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/foomo/sesamy-go/encoding/gtag"
	"github.com/pkg/errors"
)

func EventHandler(eventHandler func(event *gtag.Payload, msg *message.Message) error) func(msg *message.Message) ([]*message.Message, error) {
	return func(msg *message.Message) ([]*message.Message, error) {
		var event *gtag.Payload
		if err := json.Unmarshal(msg.Payload, &event); err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal event")
		}

		if err := eventHandler(event, msg); err != nil {
			return nil, err
		}

		b, err := json.Marshal(event)
		if err != nil {
			return nil, errors.Wrap(err, "failed to marshal event")
		}
		msg.Payload = b

		return []*message.Message{msg}, nil
	}
}
