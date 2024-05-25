package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type TutorialBegin sesamy.Event[params.TutorialBegin]

func NewTutorialBegin(p params.TutorialBegin) sesamy.Event[params.TutorialBegin] {
	return sesamy.NewEvent(sesamy.EventNameTutorialBegin, p)
}
