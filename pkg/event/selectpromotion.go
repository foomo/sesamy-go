package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type SelectPromotion sesamy.Event[params.SelectPromotion[params.Item]]

func NewSelectPromotion(p params.SelectPromotion[params.Item]) SelectPromotion {
	return SelectPromotion(sesamy.NewEvent(sesamy.EventNameSelectPromotion, p))
}
