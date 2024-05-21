package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type Refund sesamy.Event[params.Refund[params.Item]]

func NewRefund(p params.Refund[params.Item]) Refund {
	return Refund(sesamy.NewEvent(sesamy.EventNameRefund, p))
}
