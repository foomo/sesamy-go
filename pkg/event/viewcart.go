package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type ViewCart sesamy.Event[params.ViewCart[params.Item]]

func NewViewCart(p params.ViewCart[params.Item]) ViewCart {
	return ViewCart(sesamy.NewEvent(sesamy.EventNameViewCart, p))
}
