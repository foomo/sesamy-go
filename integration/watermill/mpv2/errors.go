package mpv2

import (
	"errors"
)

var (
	ErrMissingEventName = errors.New("missing event name")
	ErrErrorResponse    = errors.New("server responded with error status")
	ErrPublisherClosed  = errors.New("publisher is closed")
	ErrContextCanceled  = errors.New("request stopped without ACK received")
	ErrMessageNacked    = errors.New("message nacked")
	ErrClosed           = errors.New("subscriber already closed")
)
