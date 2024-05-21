package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type Purchase sesamy2.Event[params.Purchase[params.Item]]

func NewPurchase(p params.Purchase[params.Item]) Purchase {
	return Purchase(sesamy2.NewEvent(sesamy2.EventNamePurchase, p))
}
