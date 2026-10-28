package event

import (
	sesamyparams "github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/provider/tracify/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// EventNameTracifyAddToCart is the name of the tracify_add_to_cart event.
const EventNameTracifyAddToCart sesamy.EventName = "tracify_add_to_cart"

// TracifyAddToCart is the tracify_add_to_cart event.
type TracifyAddToCart sesamy.Event[params.TracifyAddToCart[sesamyparams.Item]]

// NewTracifyAddToCart returns a tracify_add_to_cart event with the parameters p.
func NewTracifyAddToCart(p params.TracifyAddToCart[sesamyparams.Item]) sesamy.Event[params.TracifyAddToCart[sesamyparams.Item]] {
	return sesamy.NewEvent(EventNameTracifyAddToCart, p)
}
