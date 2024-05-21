package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type BeginCheckout sesamy2.Event[params.BeginCheckout[params.Item]]

func NewBeginCheckout(p params.BeginCheckout[params.Item]) BeginCheckout {
	return BeginCheckout(sesamy2.NewEvent(sesamy2.EventNameBeginCheckout, p))
}
