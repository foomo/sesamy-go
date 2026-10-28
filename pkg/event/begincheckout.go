package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// BeginCheckout is the GA4 begin_checkout event.
type BeginCheckout sesamy.Event[params.BeginCheckout[params.Item]]

// NewBeginCheckout returns a begin_checkout event with the parameters p.
func NewBeginCheckout(p params.BeginCheckout[params.Item]) sesamy.Event[params.BeginCheckout[params.Item]] {
	return sesamy.NewEvent(sesamy.EventNameBeginCheckout, p)
}
