package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type QualifyLead sesamy.Event[params.QualifyLead[params.Item]]

func NewQualifyLead(p params.QualifyLead[params.Item]) sesamy.Event[params.QualifyLead[params.Item]] {
	return sesamy.NewEvent(sesamy.EventNameQualifyLead, p)
}
