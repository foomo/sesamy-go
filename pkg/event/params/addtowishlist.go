package params

import (
	"github.com/foomo/gostandards/iso4217"
)

// AddToWishlist https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference/events#add_to_wishlist
type AddToWishlist[I any] struct {
	Currency iso4217.Currency `json:"currency,omitempty"`
	Value    float64          `json:"value,omitempty"`
	Items    []I              `json:"items,omitempty"`
}
