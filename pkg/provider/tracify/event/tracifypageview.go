package event

import (
	"github.com/foomo/sesamy-go/pkg/provider/tracify/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// EventNameTracifyPageView is the name of the tracify_page_view event.
const EventNameTracifyPageView sesamy.EventName = "tracify_page_view"

// TracifyPageView is the tracify_page_view event.
type TracifyPageView sesamy.Event[params.TracifyPageView]

// NewTracifyPageView returns a tracify_page_view event with the parameters p.
func NewTracifyPageView(p params.TracifyPageView) sesamy.Event[params.TracifyPageView] {
	return sesamy.NewEvent(EventNameTracifyPageView, p)
}
