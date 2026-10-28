package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// ViewCart is the GA4 view_cart event.
type ViewCart sesamy.Event[params.ViewCart[params.Item]]

// NewViewCart returns a view_cart event with the parameters p.
func NewViewCart(p params.ViewCart[params.Item]) sesamy.Event[params.ViewCart[params.Item]] {
	return sesamy.NewEvent(sesamy.EventNameViewCart, p)
}
