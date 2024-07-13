package params

import (
	"github.com/foomo/gostandards/iso4217"
)

// WorkingLead
// https://developers.google.com/tag-platform/gtagjs/reference/events#working_lead
type WorkingLead[I any] struct {
	Currency   iso4217.Currency `json:"currency,omitempty"`
	Value      float64          `json:"value,omitempty"`
	Items      []I              `json:"items,omitempty"`
	LeadStatus string           `json:"lead_status,omitempty"`
}
