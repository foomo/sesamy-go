package params

import (
	"github.com/foomo/gostandards/iso4217"
)

// CloseConvertLead
// https://developers.google.com/tag-platform/gtagjs/reference/events#close_convert_lead
type CloseConvertLead[I any] struct {
	Currency iso4217.Currency `json:"currency,omitempty"`
	Value    float64          `json:"value,omitempty"`
	Items    []I              `json:"items,omitempty"`
}
