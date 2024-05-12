package params

// AddPaymentInfo https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference/events
type AddPaymentInfo[Item any] struct {
	Currency    string  `json:"currency,omitempty"`
	Value       float64 `json:"value,omitempty"`
	Coupon      string  `json:"coupon,omitempty"`
	PaymentType string  `json:"payment_type,omitempty"`
	Items       []Item  `json:"items,omitempty"`
}
