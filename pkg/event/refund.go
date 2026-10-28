package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// Refund is the GA4 refund event.
type Refund sesamy.Event[params.Refund[params.Item]]

// NewRefund returns a refund event with the parameters p.
func NewRefund(p params.Refund[params.Item]) sesamy.Event[params.Refund[params.Item]] {
	return sesamy.NewEvent(sesamy.EventNameRefund, p)
}
