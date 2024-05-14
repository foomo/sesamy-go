package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/event/params"
)

type AddPaymentInfo sesamy.Event[params.AddPaymentInfo[params.Item]]

func NewAddPaymentInfo(p params.AddPaymentInfo[params.Item]) AddPaymentInfo {
	return AddPaymentInfo(sesamy.NewEvent(sesamy.EventNameAddPaymentInfo, p))
}
