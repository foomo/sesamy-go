package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type Click sesamy2.Event[params.Click]

func NewClick(p params.Click) Click {
	return Click(sesamy2.NewEvent(sesamy2.EventNameClick, p))
}
