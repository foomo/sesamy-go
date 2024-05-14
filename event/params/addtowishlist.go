package params

// AddToWishlist https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference/events#add_to_wishlist
type AddToWishlist[Item any] struct {
	Currency string  `json:"currency,omitempty"`
	Value    float64 `json:"value,omitempty"`
	Items    []Item  `json:"items,omitempty"`
}