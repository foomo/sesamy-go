package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type FormSubmit sesamy.Event[params.FormSubmit]

func NewFormSubmit(p params.FormSubmit) FormSubmit {
	return FormSubmit(sesamy.NewEvent(sesamy.EventNameFormSubmit, p))
}
