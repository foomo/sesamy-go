package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type Refund sesamy2.Event[params.Refund[params.Item]]

func NewRefund(p params.Refund[params.Item]) Refund {
	return Refund(sesamy2.NewEvent(sesamy2.EventNameRefund, p))
}
