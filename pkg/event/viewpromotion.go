package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type ViewPromotion sesamy2.Event[params.ViewPromotion[params.Item]]

func NewViewPromotion(p params.ViewPromotion[params.Item]) ViewPromotion {
	return ViewPromotion(sesamy2.NewEvent(sesamy2.EventNameViewPromotion, p))
}
