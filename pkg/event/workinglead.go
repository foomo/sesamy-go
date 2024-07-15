package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type WorkingLead sesamy.Event[params.WorkingLead[params.Item]]

func NewWorkingLead(p params.WorkingLead[params.Item]) sesamy.Event[params.WorkingLead[params.Item]] {
	return sesamy.NewEvent(sesamy.EventNameWorkingLead, p)
}
