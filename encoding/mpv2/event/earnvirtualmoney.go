package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/encoding/mpv2/params"
)

type EarnVirtualMoney sesamy.Event[params.EarnVirtualMoney]

func NewEarnVirtualMoney(p params.EarnVirtualMoney) EarnVirtualMoney {
	return EarnVirtualMoney(sesamy.NewEvent(sesamy.EventNameEarnVirtualMoney, p))
}
