package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type AddShippingInfo sesamy2.Event[params.AddShippingInfo[params.Item]]

func NewAddShippingInfo(p params.AddShippingInfo[params.Item]) AddShippingInfo {
	return AddShippingInfo(sesamy2.NewEvent(sesamy2.EventNameAddShippingInfo, p))
}
