package params

// BeginCheckout https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference/events#begin_checkout
type BeginCheckout[Item any] struct {
	Currency string  `json:"currency,omitempty"`
	Value    float64 `json:"value,omitempty"`
	Coupon   string  `json:"coupon,omitempty"`
	Items    []Item  `json:"items,omitempty"`
}
