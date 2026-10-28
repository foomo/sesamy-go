package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// FormStart is the GA4 form_start event.
type FormStart sesamy.Event[params.FormStart]

// NewFormStart returns a form_start event with the parameters p.
func NewFormStart(p params.FormStart) sesamy.Event[params.FormStart] {
	return sesamy.NewEvent(sesamy.EventNameFormStart, p)
}
