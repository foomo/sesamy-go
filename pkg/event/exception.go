package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// Exception is the GA4 exception event.
type Exception sesamy.Event[params.Exception]

// NewException returns an exception event with the parameters p.
func NewException(p params.Exception) sesamy.Event[params.Exception] {
	return sesamy.NewEvent(sesamy.EventNameException, p)
}
