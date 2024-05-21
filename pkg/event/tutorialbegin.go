package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type TutorialBegin sesamy.Event[params.TutorialBegin]

func NewTutorialBegin(p params.TutorialBegin) TutorialBegin {
	return TutorialBegin(sesamy.NewEvent(sesamy.EventNameTutorialBegin, p))
}
