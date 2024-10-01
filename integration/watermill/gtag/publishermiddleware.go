package gtag

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"go.uber.org/zap"
)

func PublisherMiddlewareIgnoreError(next PublisherHandler) PublisherHandler {
	return func(l *zap.Logger, msg *message.Message) error {
		err := next(l, msg)
		l.With(zap.Error(err)).Warn("ignoring error")
		return nil
	}
}
