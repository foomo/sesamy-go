package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type Click sesamy.Event[params.Click]

func NewClick(p params.Click) Click {
	return Click(sesamy.NewEvent(sesamy.EventNameClick, p))
}
