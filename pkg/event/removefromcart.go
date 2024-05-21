package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type RemoveFromCart sesamy2.Event[params.RemoveFromCart[params.Item]]

func NewRemoveFromCart(p params.RemoveFromCart[params.Item]) RemoveFromCart {
	return RemoveFromCart(sesamy2.NewEvent(sesamy2.EventNameRemoveFromCart, p))
}
