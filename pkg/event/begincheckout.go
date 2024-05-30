package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type BeginCheckout sesamy.Event[params.BeginCheckout[params.Item]]

func NewBeginCheckout(p params.BeginCheckout[params.Item]) sesamy.Event[params.BeginCheckout[params.Item]] {
	return sesamy.NewEvent(sesamy.EventNameBeginCheckout, p)
}
