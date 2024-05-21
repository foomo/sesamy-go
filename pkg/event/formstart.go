package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type FormStart sesamy.Event[params.FormStart]

func NewFormStart(p params.FormStart) FormStart {
	return FormStart(sesamy.NewEvent(sesamy.EventNameFormStart, p))
}
