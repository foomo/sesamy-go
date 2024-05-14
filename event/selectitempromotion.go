package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/event/params"
)

type SelectPromotion sesamy.Event[params.SelectPromotion[params.Item]]

func NewSelectPromotion(p params.SelectPromotion[params.Item]) SelectPromotion {
	return SelectPromotion(sesamy.NewEvent(sesamy.EventNameSelectPromotion, p))
}
