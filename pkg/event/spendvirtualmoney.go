package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type SpendVirtualCurrency sesamy2.Event[params.SpendVirtualCurrency]

func NewSpendVirtualCurrency(p params.SpendVirtualCurrency) SpendVirtualCurrency {
	return SpendVirtualCurrency(sesamy2.NewEvent(sesamy2.EventNameSpendVirtualCurrency, p))
}
