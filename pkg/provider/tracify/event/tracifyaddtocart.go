package event

import (
	sesamyparams "github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/provider/tracify/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

const EventNameTracifyAddToCart sesamy.EventName = "tracify_add_to_cart"

type TracifyAddToCart sesamy.Event[params.TracifyAddToCart[sesamyparams.Item]]

func NewTracifyAddToCart(p params.TracifyAddToCart[sesamyparams.Item]) sesamy.Event[params.TracifyAddToCart[sesamyparams.Item]] {
	return sesamy.NewEvent(EventNameTracifyAddToCart, p)
}
