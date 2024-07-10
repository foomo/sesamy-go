package params

import (
	"github.com/foomo/gostandards/iso4217"
)

// AddShippingInfo https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference/events#add_shipping_info
type AddShippingInfo[I any] struct {
	Currency     iso4217.Currency `json:"currency,omitempty"`
	Value        float64          `json:"value,omitempty"`
	Coupon       string           `json:"coupon,omitempty"`
	ShippingTier string           `json:"shipping_tier,omitempty"`
	Items        []I              `json:"items,omitempty"`
}
