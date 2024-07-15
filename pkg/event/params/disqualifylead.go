package params

import (
	"github.com/foomo/gostandards/iso4217"
)

// DisqualifyLead
// https://developers.google.com/tag-platform/gtagjs/reference/events#disqualify_lead
type DisqualifyLead[I any] struct {
	Currency               iso4217.Currency `json:"currency,omitempty"`
	Value                  float64          `json:"value,omitempty"`
	Items                  []I              `json:"items,omitempty"`
	DisqualifiedLeadReason string           `json:"disqualified_lead_reason,omitempty"`
}
