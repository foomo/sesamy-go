package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type FormSubmit sesamy2.Event[params.FormSubmit]

func NewFormSubmit(p params.FormSubmit) FormSubmit {
	return FormSubmit(sesamy2.NewEvent(sesamy2.EventNameFormSubmit, p))
}
