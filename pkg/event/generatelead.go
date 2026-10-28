package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// GenerateLead is the GA4 generate_lead event.
type GenerateLead sesamy.Event[params.GenerateLead]

// NewGenerateLead returns a generate_lead event with the parameters p.
func NewGenerateLead(p params.GenerateLead) sesamy.Event[params.GenerateLead] {
	return sesamy.NewEvent(sesamy.EventNameGenerateLead, p)
}
