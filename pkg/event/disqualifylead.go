package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type DisqualifyLead sesamy.Event[params.DisqualifyLead[params.Item]]

func NewDisqualifyLead(p params.DisqualifyLead[params.Item]) sesamy.Event[params.DisqualifyLead[params.Item]] {
	return sesamy.NewEvent(sesamy.EventNameDisqualifyLead, p)
}
