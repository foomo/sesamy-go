package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// AddPaymentInfo is the GA4 add_payment_info event.
type AddPaymentInfo sesamy.Event[params.AddPaymentInfo[params.Item]]

// NewAddPaymentInfo returns an add_payment_info event with the parameters p.
func NewAddPaymentInfo(p params.AddPaymentInfo[params.Item]) sesamy.Event[params.AddPaymentInfo[params.Item]] {
	return sesamy.NewEvent(sesamy.EventNameAddPaymentInfo, p)
}
