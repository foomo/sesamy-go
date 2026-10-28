package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// SelectPromotion is the GA4 select_promotion event.
type SelectPromotion sesamy.Event[params.SelectPromotion[params.Item]]

// NewSelectPromotion returns a select_promotion event with the parameters p.
func NewSelectPromotion(p params.SelectPromotion[params.Item]) sesamy.Event[params.SelectPromotion[params.Item]] {
	return sesamy.NewEvent(sesamy.EventNameSelectPromotion, p)
}
