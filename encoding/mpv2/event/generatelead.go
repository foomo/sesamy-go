package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/encoding/mpv2/params"
)

type GenerateLead sesamy.Event[params.GenerateLead]

func NewGenerateLead(p params.GenerateLead) GenerateLead {
	return GenerateLead(sesamy.NewEvent(sesamy.EventNameGenerateLead, p))
}
