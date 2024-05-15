package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/event/params"
)

type FormStart sesamy.Event[params.FormStart]

func NewFormStart(p params.FormStart) FormStart {
	return FormStart(sesamy.NewEvent(sesamy.EventNameFormStart, p))
}
