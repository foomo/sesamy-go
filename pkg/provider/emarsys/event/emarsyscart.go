package event

import (
	sesamyparams "github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/provider/emarsys/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// EventNameEmarsysCart is the name of the emarsys_cart event.
const EventNameEmarsysCart sesamy.EventName = "emarsys_cart"

// EmarsysCart is the emarsys_cart event.
type EmarsysCart sesamy.Event[params.EmarsysCart[sesamyparams.Item]]

// NewEmarsysCart returns an emarsys_cart event with the parameters p.
func NewEmarsysCart(p params.EmarsysCart[sesamyparams.Item]) sesamy.Event[params.EmarsysCart[sesamyparams.Item]] {
	return sesamy.NewEvent(EventNameEmarsysCart, p)
}
