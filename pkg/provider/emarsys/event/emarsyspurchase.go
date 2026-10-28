package event

import (
	sesamyparams "github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/provider/emarsys/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// EventNameEmarsysPurchase is the name of the emarsys_purchase event.
const EventNameEmarsysPurchase sesamy.EventName = "emarsys_purchase"

// EmarsysPurchase is the emarsys_purchase event.
type EmarsysPurchase sesamy.Event[params.EmarsysPurchase[sesamyparams.Item]]

// NewEmarsysPurchase returns an emarsys_purchase event with the parameters p.
func NewEmarsysPurchase(p params.EmarsysPurchase[sesamyparams.Item]) sesamy.Event[params.EmarsysPurchase[sesamyparams.Item]] {
	return sesamy.NewEvent(EventNameEmarsysPurchase, p)
}
