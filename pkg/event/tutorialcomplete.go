package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type TutorialComplete sesamy2.Event[params.TutorialComplete]

func NewTutorialComplete(p params.TutorialComplete) TutorialComplete {
	return TutorialComplete(sesamy2.NewEvent(sesamy2.EventNameTutorialComplete, p))
}
