package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type AdImpression sesamy.Event[params.AdImpression]

func NewAdImpression(p params.AdImpression) AdImpression {
	return AdImpression(sesamy.NewEvent(sesamy.EventNameAdImpression, p))
}
