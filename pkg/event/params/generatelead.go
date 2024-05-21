package params

import (
	"github.com/foomo/gostandards/iso4217"
)

// GenerateLead https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference/events#generate_lead
type GenerateLead struct {
	Currency iso4217.Currency `json:"currency,omitempty"`
	Value    float64          `json:"value,omitempty"`
}
