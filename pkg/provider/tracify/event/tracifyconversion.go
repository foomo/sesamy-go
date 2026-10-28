package event

import (
	"github.com/foomo/sesamy-go/pkg/provider/tracify/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// EventNameTracifyConversion is the name of the tracify_conversion event.
const EventNameTracifyConversion sesamy.EventName = "tracify_conversion"

// TracifyConversion is the tracify_conversion event.
type TracifyConversion sesamy.Event[params.TracifyConversion]

// NewTracifyConversion returns a tracify_conversion event with the parameters p.
func NewTracifyConversion(p params.TracifyConversion) sesamy.Event[params.TracifyConversion] {
	return sesamy.NewEvent(EventNameTracifyConversion, p)
}
