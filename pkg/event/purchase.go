package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type Purchase sesamy.Event[params.Purchase[params.Item]]

func NewPurchase(p params.Purchase[params.Item]) Purchase {
	return Purchase(sesamy.NewEvent(sesamy.EventNamePurchase, p))
}
