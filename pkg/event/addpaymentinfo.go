package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type AddPaymentInfo sesamy.Event[params.AddPaymentInfo[params.Item]]

func NewAddPaymentInfo(p params.AddPaymentInfo[params.Item]) AddPaymentInfo {
	return AddPaymentInfo(sesamy.NewEvent(sesamy.EventNameAddPaymentInfo, p))
}
