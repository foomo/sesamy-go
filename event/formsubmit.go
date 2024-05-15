package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/event/params"
)

type FormSubmit sesamy.Event[params.FormSubmit]

func NewFormSubmit(p params.FormSubmit) FormSubmit {
	return FormSubmit(sesamy.NewEvent(sesamy.EventNameFormSubmit, p))
}
