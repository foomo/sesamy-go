package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type SelectPromotion sesamy2.Event[params.SelectPromotion[params.Item]]

func NewSelectPromotion(p params.SelectPromotion[params.Item]) SelectPromotion {
	return SelectPromotion(sesamy2.NewEvent(sesamy2.EventNameSelectPromotion, p))
}
