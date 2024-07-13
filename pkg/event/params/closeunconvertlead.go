package params

import (
	"github.com/foomo/gostandards/iso4217"
)

// CloseUnconvertLead
// https://developers.google.com/tag-platform/gtagjs/reference/events#close_unconvert_lead
type CloseUnconvertLead[I any] struct {
	Currency            iso4217.Currency `json:"currency,omitempty"`
	Value               float64          `json:"value,omitempty"`
	Items               []I              `json:"items,omitempty"`
	UnconvertLeadReason string           `json:"unconvert_lead_reason,omitempty"`
}
