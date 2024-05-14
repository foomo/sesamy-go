package params

// ViewCart https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference/events#view_cart
type ViewCart[Item any] struct {
	Currency string  `json:"currency,omitempty"`
	Value    float64 `json:"value,omitempty"`
	Items    []Item  `json:"items,omitempty"`
}
