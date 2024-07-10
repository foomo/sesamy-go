package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type GenerateLead sesamy.Event[params.GenerateLead]

func NewGenerateLead(p params.GenerateLead) sesamy.Event[params.GenerateLead] {
	return sesamy.NewEvent(sesamy.EventNameGenerateLead, p)
}
