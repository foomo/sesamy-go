package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type CloseConvertLead sesamy.Event[params.CloseConvertLead[params.Item]]

func NewCloseConvertLead(p params.CloseConvertLead[params.Item]) sesamy.Event[params.CloseConvertLead[params.Item]] {
	return sesamy.NewEvent(sesamy.EventNameCloseConvertLead, p)
}
