package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type TutorialComplete sesamy.Event[params.TutorialComplete]

func NewTutorialComplete(p params.TutorialComplete) sesamy.Event[params.TutorialComplete] {
	return sesamy.NewEvent(sesamy.EventNameTutorialComplete, p)
}
