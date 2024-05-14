package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/event/params"
)

type RemoveFromCart sesamy.Event[params.RemoveFromCart[params.Item]]

func NewRemoveFromCart(p params.RemoveFromCart[params.Item]) RemoveFromCart {
	return RemoveFromCart(sesamy.NewEvent(sesamy.EventNameRemoveFromCart, p))
}