package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// Purchase is the GA4 purchase event.
type Purchase sesamy.Event[params.Purchase[params.Item]]

// NewPurchase returns a purchase event with the parameters p.
func NewPurchase(p params.Purchase[params.Item]) sesamy.Event[params.Purchase[params.Item]] {
	return sesamy.NewEvent(sesamy.EventNamePurchase, p)
}
