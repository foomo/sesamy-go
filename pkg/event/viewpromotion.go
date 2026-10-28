package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// ViewPromotion is the GA4 view_promotion event.
type ViewPromotion sesamy.Event[params.ViewPromotion[params.Item]]

// NewViewPromotion returns a view_promotion event with the parameters p.
func NewViewPromotion(p params.ViewPromotion[params.Item]) sesamy.Event[params.ViewPromotion[params.Item]] {
	return sesamy.NewEvent(sesamy.EventNameViewPromotion, p)
}
