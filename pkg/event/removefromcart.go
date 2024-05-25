package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type RemoveFromCart sesamy.Event[params.RemoveFromCart[params.Item]]

func NewRemoveFromCart(p params.RemoveFromCart[params.Item]) sesamy.Event[params.RemoveFromCart[params.Item]] {
	return sesamy.NewEvent(sesamy.EventNameRemoveFromCart, p)
}
