package params

import (
	"github.com/foomo/gostandards/iso4217"
)

// TracifyConversion holds the parameters of the tracify_conversion event.
type TracifyConversion struct {
	Currency     iso4217.Code `json:"currency,omitempty"`
	Value        float64      `json:"value,omitempty"`
	ConversionID string       `json:"conversion_id,omitempty"`
}
