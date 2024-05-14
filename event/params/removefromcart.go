package params

// RemoveFromCart https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference/events#remove_from_cart
type RemoveFromCart[Item any] struct {
	Currency string  `json:"currency,omitempty"`
	Value    float64 `json:"value,omitempty"`
	Items    []Item  `json:"items,omitempty"`
}
