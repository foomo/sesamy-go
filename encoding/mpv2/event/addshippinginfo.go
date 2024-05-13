package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/encoding/mpv2/params"
)

type AddShippingInfo sesamy.Event[params.AddShippingInfo[params.Item]]

func NewAddShippingInfo(p params.AddShippingInfo[params.Item]) AddShippingInfo {
	return AddShippingInfo(sesamy.NewEvent(sesamy.EventNameAddShippingInfo, p))
}
