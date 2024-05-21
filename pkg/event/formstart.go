package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type FormStart sesamy2.Event[params.FormStart]

func NewFormStart(p params.FormStart) FormStart {
	return FormStart(sesamy2.NewEvent(sesamy2.EventNameFormStart, p))
}
