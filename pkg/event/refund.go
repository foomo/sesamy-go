package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type Refund sesamy.Event[params.Refund[params.Item]]

func NewRefund(p params.Refund[params.Item]) sesamy.Event[params.Refund[params.Item]] {
	return sesamy.NewEvent(sesamy.EventNameRefund, p)
}
