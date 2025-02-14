package event

import (
	sesamyparams "github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/provider/emarsys/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

const EventNameEmarsysAddToCart sesamy.EventName = "emarsys_add_to_cart"

type EmarsysAddToCart sesamy.Event[params.EmarsysAddToCart[sesamyparams.Item]]

func NewEmarsysAddToCart(p params.EmarsysAddToCart[sesamyparams.Item]) sesamy.Event[params.EmarsysAddToCart[sesamyparams.Item]] {
	return sesamy.NewEvent(EventNameEmarsysAddToCart, p)
}
