package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type EarnVirtualMoney sesamy.Event[params.EarnVirtualMoney]

func NewEarnVirtualMoney(p params.EarnVirtualMoney) EarnVirtualMoney {
	return EarnVirtualMoney(sesamy.NewEvent(sesamy.EventNameEarnVirtualMoney, p))
}
