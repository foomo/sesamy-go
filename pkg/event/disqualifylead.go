package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// DisqualifyLead is the GA4 disqualify_lead event.
type DisqualifyLead sesamy.Event[params.DisqualifyLead[params.Item]]

// NewDisqualifyLead returns a disqualify_lead event with the parameters p.
func NewDisqualifyLead(p params.DisqualifyLead[params.Item]) sesamy.Event[params.DisqualifyLead[params.Item]] {
	return sesamy.NewEvent(sesamy.EventNameDisqualifyLead, p)
}
