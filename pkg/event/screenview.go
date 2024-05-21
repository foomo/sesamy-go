package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type ScreenView sesamy2.Event[params.ScreenView]

func NewScreenView(p params.ScreenView) ScreenView {
	return ScreenView(sesamy2.NewEvent(sesamy2.EventNameScreenView, p))
}
