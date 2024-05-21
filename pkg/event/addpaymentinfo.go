package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type AddPaymentInfo sesamy2.Event[params.AddPaymentInfo[params.Item]]

func NewAddPaymentInfo(p params.AddPaymentInfo[params.Item]) AddPaymentInfo {
	return AddPaymentInfo(sesamy2.NewEvent(sesamy2.EventNameAddPaymentInfo, p))
}
