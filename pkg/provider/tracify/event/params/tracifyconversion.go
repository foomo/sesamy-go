package params

import (
	"github.com/foomo/gostandards/iso4217"
)

type TracifyConversion struct {
	Currency     iso4217.Currency `json:"currency,omitempty"`
	Value        float64          `json:"value,omitempty"`
	ConversionID string           `json:"conversion_id,omitempty"`
}
