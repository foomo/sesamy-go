package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/event/params"
)

type ScreenView sesamy.Event[params.ScreenView]

func NewScreenView(p params.ScreenView) ScreenView {
	return ScreenView(sesamy.NewEvent(sesamy.EventNameScreenView, p))
}
