package params

import (
	"github.com/foomo/gostandards/iso4217"
)

// ViewCart https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference/events#view_cart
type ViewCart[Item any] struct {
	Currency iso4217.Currency `json:"currency,omitempty"`
	Value    float64          `json:"value,omitempty"`
	Items    []Item           `json:"items,omitempty"`
}
