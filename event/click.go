package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/event/params"
)

type Click sesamy.Event[params.Click]

func NewClick(p params.Click) Click {
	return Click(sesamy.NewEvent(sesamy.EventNameClick, p))
}
