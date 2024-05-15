package params

import (
	"github.com/foomo/gostandards/iso4217"
)

// Purchase https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference/events#purchase
type Purchase[Item any] struct {
	Currency      iso4217.Currency `json:"currency,omitempty"`
	Value         float64          `json:"value,omitempty"`
	TransactionID string           `json:"transaction_id,omitempty"`
	Coupon        string           `json:"coupon,omitempty"`
	Shipping      float64          `json:"shipping,omitempty"`
	Tax           float64          `json:"tax,omitempty"`
	Items         []Item           `json:"items,omitempty"`
}
