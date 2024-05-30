package params

import (
	"github.com/foomo/gostandards/iso4217"
)

// AddPaymentInfo https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference/events
type AddPaymentInfo[I any] struct {
	Currency    iso4217.Currency `json:"currency,omitempty"`
	Value       float64          `json:"value,omitempty"`
	Coupon      string           `json:"coupon,omitempty"`
	PaymentType string           `json:"payment_type,omitempty"`
	Items       []I              `json:"items,omitempty"`
}
