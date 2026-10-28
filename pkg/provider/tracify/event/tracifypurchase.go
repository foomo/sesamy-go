package event

import (
	sesamyparams "github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/provider/tracify/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// EventNameTracifyPurchase is the name of the tracify_purchase event.
const EventNameTracifyPurchase sesamy.EventName = "tracify_purchase"

// TracifyPurchase is the tracify_purchase event.
type TracifyPurchase sesamy.Event[params.TracifyPurchase[sesamyparams.Item]]

// NewTracifyPurchase returns a tracify_purchase event with the parameters p.
func NewTracifyPurchase(p params.TracifyPurchase[sesamyparams.Item]) sesamy.Event[params.TracifyPurchase[sesamyparams.Item]] {
	return sesamy.NewEvent(EventNameTracifyPurchase, p)
}
