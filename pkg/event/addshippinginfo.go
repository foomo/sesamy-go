package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type AddShippingInfo sesamy.Event[params.AddShippingInfo[params.Item]]

func NewAddShippingInfo(p params.AddShippingInfo[params.Item]) sesamy.Event[params.AddShippingInfo[params.Item]] {
	return sesamy.NewEvent(sesamy.EventNameAddShippingInfo, p)
}
