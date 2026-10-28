package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// SpendVirtualCurrency is the GA4 spend_virtual_currency event.
type SpendVirtualCurrency sesamy.Event[params.SpendVirtualCurrency]

// NewSpendVirtualCurrency returns a spend_virtual_currency event with the parameters p.
func NewSpendVirtualCurrency(p params.SpendVirtualCurrency) sesamy.Event[params.SpendVirtualCurrency] {
	return sesamy.NewEvent(sesamy.EventNameSpendVirtualCurrency, p)
}
