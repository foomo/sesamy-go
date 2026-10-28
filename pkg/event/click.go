package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// Click is the GA4 click event.
type Click sesamy.Event[params.Click]

// NewClick returns a click event with the parameters p.
func NewClick(p params.Click) sesamy.Event[params.Click] {
	return sesamy.NewEvent(sesamy.EventNameClick, p)
}
