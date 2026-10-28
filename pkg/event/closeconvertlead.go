package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// CloseConvertLead is the GA4 close_convert_lead event.
type CloseConvertLead sesamy.Event[params.CloseConvertLead[params.Item]]

// NewCloseConvertLead returns a close_convert_lead event with the parameters p.
func NewCloseConvertLead(p params.CloseConvertLead[params.Item]) sesamy.Event[params.CloseConvertLead[params.Item]] {
	return sesamy.NewEvent(sesamy.EventNameCloseConvertLead, p)
}
