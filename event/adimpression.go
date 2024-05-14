package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/event/params"
)

type AdImpression sesamy.Event[params.AdImpression]

func NewAdImpression(p params.AdImpression) AdImpression {
	return AdImpression(sesamy.NewEvent(sesamy.EventNameAdImpression, p))
}