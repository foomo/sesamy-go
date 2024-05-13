package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/encoding/mpv2/params"
)

type AddToCart sesamy.Event[params.AddToCart[params.Item]]

func NewAddToCart(p params.AddToCart[params.Item]) AddToCart {
	return AddToCart(sesamy.NewEvent(sesamy.EventNameAddToCart, p))
}
