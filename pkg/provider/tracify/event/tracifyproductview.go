package event

import (
	"github.com/foomo/sesamy-go/pkg/provider/tracify/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// EventNameTracifyProductView is the name of the tracify_product_view event.
const EventNameTracifyProductView sesamy.EventName = "tracify_product_view"

// TracifyProductView is the tracify_product_view event.
type TracifyProductView sesamy.Event[params.TracifyProductView]

// NewTracifyProductView returns a tracify_product_view event with the parameters p.
func NewTracifyProductView(p params.TracifyProductView) sesamy.Event[params.TracifyProductView] {
	return sesamy.NewEvent(EventNameTracifyProductView, p)
}
