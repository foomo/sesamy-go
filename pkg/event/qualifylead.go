package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// QualifyLead is the GA4 qualify_lead event.
type QualifyLead sesamy.Event[params.QualifyLead[params.Item]]

// NewQualifyLead returns a qualify_lead event with the parameters p.
func NewQualifyLead(p params.QualifyLead[params.Item]) sesamy.Event[params.QualifyLead[params.Item]] {
	return sesamy.NewEvent(sesamy.EventNameQualifyLead, p)
}
