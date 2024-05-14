package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/event/params"
)

type BeginCheckout sesamy.Event[params.BeginCheckout[params.Item]]

func NewBeginCheckout(p params.BeginCheckout[params.Item]) BeginCheckout {
	return BeginCheckout(sesamy.NewEvent(sesamy.EventNameBeginCheckout, p))
}
