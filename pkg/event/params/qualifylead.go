package params

import (
	"github.com/foomo/gostandards/iso4217"
)

// QualifyLead
// https://developers.google.com/tag-platform/gtagjs/reference/events#qualify_lead
type QualifyLead[I any] struct {
	Currency iso4217.Currency `json:"currency,omitempty"`
	Value    float64          `json:"value,omitempty"`
	Items    []I              `json:"items,omitempty"`
}
