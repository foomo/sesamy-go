package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type GenerateLead sesamy2.Event[params.GenerateLead]

func NewGenerateLead(p params.GenerateLead) GenerateLead {
	return GenerateLead(sesamy2.NewEvent(sesamy2.EventNameGenerateLead, p))
}
