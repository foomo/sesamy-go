package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type ViewCart sesamy.Event[params.ViewCart[params.Item]]

func NewViewCart(p params.ViewCart[params.Item]) sesamy.Event[params.ViewCart[params.Item]] {
	return sesamy.NewEvent(sesamy.EventNameViewCart, p)
}
