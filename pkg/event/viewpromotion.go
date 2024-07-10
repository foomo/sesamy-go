package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type ViewPromotion sesamy.Event[params.ViewPromotion[params.Item]]

func NewViewPromotion(p params.ViewPromotion[params.Item]) sesamy.Event[params.ViewPromotion[params.Item]] {
	return sesamy.NewEvent(sesamy.EventNameViewPromotion, p)
}
