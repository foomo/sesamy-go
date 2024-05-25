package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type SpendVirtualCurrency sesamy.Event[params.SpendVirtualCurrency]

func NewSpendVirtualCurrency(p params.SpendVirtualCurrency) sesamy.Event[params.SpendVirtualCurrency] {
	return sesamy.NewEvent(sesamy.EventNameSpendVirtualCurrency, p)
}
