package params

import (
	"github.com/foomo/gostandards/iso4217"
)

// TracifyAddToCart holds the parameters of the tracify_add_to_cart event.
type TracifyAddToCart[I any] struct {
	Currency iso4217.Code `json:"currency,omitempty"`
	Value    float64      `json:"value,omitempty"`
	Items    []I          `json:"items,omitempty"`
}
