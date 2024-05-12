package params

// GenerateLead https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference/events#generate_lead
type GenerateLead struct {
	Currency string  `json:"currency,omitempty"`
	Value    float64 `json:"value,omitempty"`
}
