package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// FormSubmit is the GA4 form_submit event.
type FormSubmit sesamy.Event[params.FormSubmit]

// NewFormSubmit returns a form_submit event with the parameters p.
func NewFormSubmit(p params.FormSubmit) sesamy.Event[params.FormSubmit] {
	return sesamy.NewEvent(sesamy.EventNameFormSubmit, p)
}
