package event

import (
	"github.com/foomo/sesamy-go/pkg/provider/tracify/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

const EventNameTracifyPageView sesamy.EventName = "tracify_page_view"

type TracifyPageView sesamy.Event[params.TracifyPageView]

func NewTracifyPageView(p params.TracifyPageView) sesamy.Event[params.TracifyPageView] {
	return sesamy.NewEvent(EventNameTracifyPageView, p)
}
