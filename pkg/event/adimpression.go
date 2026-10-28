package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// AdImpression is the GA4 ad_impression event.
type AdImpression sesamy.Event[params.AdImpression]

// NewAdImpression returns an ad_impression event with the parameters p.
func NewAdImpression(p params.AdImpression) sesamy.Event[params.AdImpression] {
	return sesamy.NewEvent(sesamy.EventNameAdImpression, p)
}
