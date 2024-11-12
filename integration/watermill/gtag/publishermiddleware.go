package gtag

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

func PublisherMiddlewareIgnoreError(next PublisherHandler) PublisherHandler {
	return func(l *zap.Logger, msg *message.Message) error {
		if err := next(l, msg); err != nil {
			if spanCtx := trace.SpanContextFromContext(msg.Context()); spanCtx.IsValid() && spanCtx.IsSampled() {
				l = l.With(zap.String("trace_id", spanCtx.TraceID().String()), zap.String("span_id", spanCtx.SpanID().String()))
			}
			l.With(zap.Error(err)).Warn("ignoring error")
		}
		return nil
	}
}
