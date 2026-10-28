package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// ScreenView is the GA4 screen_view event.
type ScreenView sesamy.Event[params.ScreenView]

// NewScreenView returns a screen_view event with the parameters p.
func NewScreenView(p params.ScreenView) sesamy.Event[params.ScreenView] {
	return sesamy.NewEvent(sesamy.EventNameScreenView, p)
}
