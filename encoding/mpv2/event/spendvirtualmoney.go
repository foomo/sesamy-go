package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/encoding/mpv2/params"
)

type SpendVirtualCurrency sesamy.Event[params.SpendVirtualCurrency]

func NewSpendVirtualCurrency(p params.SpendVirtualCurrency) SpendVirtualCurrency {
	return SpendVirtualCurrency(sesamy.NewEvent(sesamy.EventNameSpendVirtualCurrency, p))
}
