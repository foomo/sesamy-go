package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// WorkingLead is the GA4 working_lead event.
type WorkingLead sesamy.Event[params.WorkingLead[params.Item]]

// NewWorkingLead returns a working_lead event with the parameters p.
func NewWorkingLead(p params.WorkingLead[params.Item]) sesamy.Event[params.WorkingLead[params.Item]] {
	return sesamy.NewEvent(sesamy.EventNameWorkingLead, p)
}
