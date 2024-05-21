package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type EarnVirtualMoney sesamy2.Event[params.EarnVirtualMoney]

func NewEarnVirtualMoney(p params.EarnVirtualMoney) EarnVirtualMoney {
	return EarnVirtualMoney(sesamy2.NewEvent(sesamy2.EventNameEarnVirtualMoney, p))
}
