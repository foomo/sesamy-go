package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/event/params"
)

type TutorialBegin sesamy.Event[params.TutorialBegin]

func NewTutorialBegin(p params.TutorialBegin) TutorialBegin {
	return TutorialBegin(sesamy.NewEvent(sesamy.EventNameTutorialBegin, p))
}
