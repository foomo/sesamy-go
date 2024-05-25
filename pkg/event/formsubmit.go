package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type FormSubmit sesamy.Event[params.FormSubmit]

func NewFormSubmit(p params.FormSubmit) sesamy.Event[params.FormSubmit] {
	return sesamy.NewEvent(sesamy.EventNameFormSubmit, p)
}
