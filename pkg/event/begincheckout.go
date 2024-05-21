package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type BeginCheckout sesamy.Event[params.BeginCheckout[params.Item]]

func NewBeginCheckout(p params.BeginCheckout[params.Item]) BeginCheckout {
	return BeginCheckout(sesamy.NewEvent(sesamy.EventNameBeginCheckout, p))
}
