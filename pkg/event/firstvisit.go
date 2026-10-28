package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// FirstVisit is the GA4 first_visit event.
type FirstVisit sesamy.Event[params.FirstVisit]

// NewFirstVisit returns a first_visit event with the parameters p.
func NewFirstVisit(p params.FirstVisit) sesamy.Event[params.FirstVisit] {
	return sesamy.NewEvent(sesamy.EventNameFirstVisit, p)
}
