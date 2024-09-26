package event

import (
	"github.com/foomo/sesamy-go/pkg/provider/tracify/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

const EventNameTracifyProductView sesamy.EventName = "tracify_product_view"

type TracifyProductView sesamy.Event[params.TracifyProductView]

func NewTracifyProductView(p params.TracifyProductView) sesamy.Event[params.TracifyProductView] {
	return sesamy.NewEvent(EventNameTracifyProductView, p)
}
