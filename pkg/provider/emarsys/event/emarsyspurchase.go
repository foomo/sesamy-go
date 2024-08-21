package event

import (
	sesamyparams "github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/provider/emarsys/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

const EventNameEmarsysPurchase sesamy.EventName = "emarsys_cart"

type EmarsysPurchase sesamy.Event[params.EmarsysPurchase[sesamyparams.Item]]

func NewEmarsysPurchase(p params.EmarsysPurchase[sesamyparams.Item]) sesamy.Event[params.EmarsysPurchase[sesamyparams.Item]] {
	return sesamy.NewEvent(EventNameEmarsysPurchase, p)
}
