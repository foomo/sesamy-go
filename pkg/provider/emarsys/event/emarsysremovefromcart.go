package event

import (
	sesamyparams "github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/provider/emarsys/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

const EventNameEmarsysRemoveFromCart sesamy.EventName = "emarsys_remove_from_cart"

type EmarsysRemoveFromCart sesamy.Event[params.EmarsysRemoveFromCart[sesamyparams.Item]]

func NewEmarsysRemoveFromCart(p params.EmarsysRemoveFromCart[sesamyparams.Item]) sesamy.Event[params.EmarsysRemoveFromCart[sesamyparams.Item]] {
	return sesamy.NewEvent(EventNameEmarsysRemoveFromCart, p)
}
