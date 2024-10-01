package params

import (
	"github.com/foomo/gostandards/iso4217"
)

type TracifyPurchase[I any] struct {
	Currency      iso4217.Currency `json:"currency,omitempty"`
	Value         float64          `json:"value,omitempty"`
	TransactionID string           `json:"transaction_id,omitempty"`
	Items         []I              `json:"items,omitempty"`
}
