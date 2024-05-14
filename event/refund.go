package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/event/params"
)

type Refund sesamy.Event[params.Refund[params.Item]]

func NewRefund(p params.Refund[params.Item]) Refund {
	return Refund(sesamy.NewEvent(sesamy.EventNameRefund, p))
}
