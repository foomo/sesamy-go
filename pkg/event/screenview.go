package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type ScreenView sesamy.Event[params.ScreenView]

func NewScreenView(p params.ScreenView) sesamy.Event[params.ScreenView] {
	return sesamy.NewEvent(sesamy.EventNameScreenView, p)
}
