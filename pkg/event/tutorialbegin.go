package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type TutorialBegin sesamy2.Event[params.TutorialBegin]

func NewTutorialBegin(p params.TutorialBegin) TutorialBegin {
	return TutorialBegin(sesamy2.NewEvent(sesamy2.EventNameTutorialBegin, p))
}
