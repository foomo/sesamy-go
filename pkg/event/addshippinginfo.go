package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// AddShippingInfo is the GA4 add_shipping_info event.
type AddShippingInfo sesamy.Event[params.AddShippingInfo[params.Item]]

// NewAddShippingInfo returns an add_shipping_info event with the parameters p.
func NewAddShippingInfo(p params.AddShippingInfo[params.Item]) sesamy.Event[params.AddShippingInfo[params.Item]] {
	return sesamy.NewEvent(sesamy.EventNameAddShippingInfo, p)
}
