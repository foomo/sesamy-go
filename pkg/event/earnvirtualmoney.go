package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// EarnVirtualMoney is the GA4 earn_virtual_money event.
type EarnVirtualMoney sesamy.Event[params.EarnVirtualMoney]

// NewEarnVirtualMoney returns an earn_virtual_money event with the parameters p.
func NewEarnVirtualMoney(p params.EarnVirtualMoney) sesamy.Event[params.EarnVirtualMoney] {
	return sesamy.NewEvent(sesamy.EventNameEarnVirtualMoney, p)
}
