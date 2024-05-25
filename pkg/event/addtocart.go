package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type AddToCart sesamy.Event[params.AddToCart[params.Item]]

func NewAddToCart(p params.AddToCart[params.Item]) sesamy.Event[params.AddToCart[params.Item]] {
	return sesamy.NewEvent(sesamy.EventNameAddToCart, p)
}
