package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type Click sesamy.Event[params.Click]

func NewClick(p params.Click) sesamy.Event[params.Click] {
	return sesamy.NewEvent(sesamy.EventNameClick, p)
}
