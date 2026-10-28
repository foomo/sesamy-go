package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// AddToCart is the GA4 add_to_cart event.
type AddToCart sesamy.Event[params.AddToCart[params.Item]]

// NewAddToCart returns an add_to_cart event with the parameters p.
func NewAddToCart(p params.AddToCart[params.Item]) sesamy.Event[params.AddToCart[params.Item]] {
	return sesamy.NewEvent(sesamy.EventNameAddToCart, p)
}
