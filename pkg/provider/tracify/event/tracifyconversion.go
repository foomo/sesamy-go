package event

import (
	"github.com/foomo/sesamy-go/pkg/provider/tracify/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

const EventNameTracifyConversion sesamy.EventName = "tracify_conversion"

type TracifyConversion sesamy.Event[params.TracifyConversion]

func NewTracifyConversion(p params.TracifyConversion) sesamy.Event[params.TracifyConversion] {
	return sesamy.NewEvent(EventNameTracifyConversion, p)
}
