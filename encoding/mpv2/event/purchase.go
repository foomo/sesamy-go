package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/encoding/mpv2/params"
)

type Purchase sesamy.Event[params.Purchase[params.Item]]

func NewPurchase(p params.Purchase[params.Item]) Purchase {
	return Purchase(sesamy.NewEvent(sesamy.EventNamePurchase, p))
}
