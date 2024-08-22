package event

import (
	sesamyparams "github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/provider/emarsys/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

const EventNameEmarsysCart sesamy.EventName = "emarsys_cart"

type EmarsysCart sesamy.Event[params.EmarsysCart[sesamyparams.Item]]

func NewEmarsysCart(p params.EmarsysCart[sesamyparams.Item]) sesamy.Event[params.EmarsysCart[sesamyparams.Item]] {
	return sesamy.NewEvent(EventNameEmarsysCart, p)
}
