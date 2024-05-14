package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/event/params"
)

type ViewPromotion sesamy.Event[params.ViewPromotion[params.Item]]

func NewViewPromotion(p params.ViewPromotion[params.Item]) ViewPromotion {
	return ViewPromotion(sesamy.NewEvent(sesamy.EventNameViewPromotion, p))
}
