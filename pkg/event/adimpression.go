package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type AdImpression sesamy2.Event[params.AdImpression]

func NewAdImpression(p params.AdImpression) AdImpression {
	return AdImpression(sesamy2.NewEvent(sesamy2.EventNameAdImpression, p))
}
