package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// CloseUnconvertLead is the GA4 close_unconvert_lead event.
type CloseUnconvertLead sesamy.Event[params.CloseUnconvertLead[params.Item]]

// NewCloseUnconvertLead returns a close_unconvert_lead event with the parameters p.
func NewCloseUnconvertLead(p params.CloseUnconvertLead[params.Item]) sesamy.Event[params.CloseUnconvertLead[params.Item]] {
	return sesamy.NewEvent(sesamy.EventNameCloseUnconvertLead, p)
}
