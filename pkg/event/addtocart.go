package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type AddToCart sesamy2.Event[params.AddToCart[params.Item]]

func NewAddToCart(p params.AddToCart[params.Item]) AddToCart {
	return AddToCart(sesamy2.NewEvent(sesamy2.EventNameAddToCart, p))
}
