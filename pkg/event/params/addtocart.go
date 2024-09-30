package params

import (
	"github.com/foomo/gostandards/iso4217"
)

// AddToCart https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference/events#add_to_cart
type AddToCart[I any] struct {
	*MPv2    `json:",inline"`
	Currency iso4217.Currency `json:"currency,omitempty"`
	Value    float64          `json:"value,omitempty"`
	Items    []I              `json:"items,omitempty"`
}
