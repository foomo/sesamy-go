package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// RemoveFromCart is the GA4 remove_from_cart event.
type RemoveFromCart sesamy.Event[params.RemoveFromCart[params.Item]]

// NewRemoveFromCart returns a remove_from_cart event with the parameters p.
func NewRemoveFromCart(p params.RemoveFromCart[params.Item]) sesamy.Event[params.RemoveFromCart[params.Item]] {
	return sesamy.NewEvent(sesamy.EventNameRemoveFromCart, p)
}
