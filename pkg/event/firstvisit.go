package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type FirstVisit sesamy.Event[params.FirstVisit]

func NewFirstVisit(p params.FirstVisit) sesamy.Event[params.FirstVisit] {
	return sesamy.NewEvent(sesamy.EventNameFirstVisit, p)
}
