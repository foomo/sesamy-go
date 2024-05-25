package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type FormStart sesamy.Event[params.FormStart]

func NewFormStart(p params.FormStart) sesamy.Event[params.FormStart] {
	return sesamy.NewEvent(sesamy.EventNameFormStart, p)
}
