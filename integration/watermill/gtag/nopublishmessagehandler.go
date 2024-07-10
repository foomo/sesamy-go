package gtag

import (
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/foomo/sesamy-go/pkg/encoding/gtag"
	"github.com/pkg/errors"
)

func NoPublishMessageHandler(handler func(payload *gtag.Payload, msg *message.Message) error) func(msg *message.Message) error {
	return func(msg *message.Message) error {
		var payload *gtag.Payload

		// unmarshal payload
		if err := json.Unmarshal(msg.Payload, &payload); err != nil {
			return errors.Wrap(err, "failed to unmarshal payload")
		}

		// handle payload
		return handler(payload, msg)
	}
}
