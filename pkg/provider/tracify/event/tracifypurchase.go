package event

import (
	sesamyparams "github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/provider/tracify/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

const EventNameTracifyPurchase sesamy.EventName = "tracify_purchase"

type TracifyPurchase sesamy.Event[params.TracifyPurchase[sesamyparams.Item]]

func NewTracifyPurchase(p params.TracifyPurchase[sesamyparams.Item]) sesamy.Event[params.TracifyPurchase[sesamyparams.Item]] {
	return sesamy.NewEvent(EventNameTracifyPurchase, p)
}
