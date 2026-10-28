package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// Scroll is the GA4 scroll event.
type Scroll sesamy.Event[params.Scroll]

// NewScroll returns a scroll event with the parameters p.
func NewScroll(p params.Scroll) sesamy.Event[params.Scroll] {
	return sesamy.NewEvent(sesamy.EventNameScroll, p)
}
