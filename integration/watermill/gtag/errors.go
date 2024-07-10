package gtag

import (
	"errors"
)

var (
	ErrMissingEventName = errors.New("missing event name")
	ErrContextCanceled  = errors.New("request stopped without ACK received")
	ErrMessageNacked    = errors.New("message nacked")
	ErrClosed           = errors.New("subscriber already closed")
)
