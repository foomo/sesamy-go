package params

import (
	"github.com/foomo/gostandards/iso4217"
)

// RemoveFromCart https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference/events#remove_from_cart
type RemoveFromCart[I any] struct {
	Currency iso4217.Currency `json:"currency,omitempty"`
	Value    float64          `json:"value,omitempty"`
	Items    []I              `json:"items,omitempty"`
}
