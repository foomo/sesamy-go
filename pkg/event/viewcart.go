package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type ViewCart sesamy2.Event[params.ViewCart[params.Item]]

func NewViewCart(p params.ViewCart[params.Item]) ViewCart {
	return ViewCart(sesamy2.NewEvent(sesamy2.EventNameViewCart, p))
}
