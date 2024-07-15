package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type CloseUnconvertLead sesamy.Event[params.CloseUnconvertLead[params.Item]]

func NewCloseUnconvertLead(p params.CloseUnconvertLead[params.Item]) sesamy.Event[params.CloseUnconvertLead[params.Item]] {
	return sesamy.NewEvent(sesamy.EventNameCloseUnconvertLead, p)
}
