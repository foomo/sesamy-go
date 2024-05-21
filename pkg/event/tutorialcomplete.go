package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type TutorialComplete sesamy.Event[params.TutorialComplete]

func NewTutorialComplete(p params.TutorialComplete) TutorialComplete {
	return TutorialComplete(sesamy.NewEvent(sesamy.EventNameTutorialComplete, p))
}
