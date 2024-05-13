package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/encoding/mpv2/params"
)

type TutorialComplete sesamy.Event[params.TutorialComplete]

func NewTutorialComplete(p params.TutorialComplete) TutorialComplete {
	return TutorialComplete(sesamy.NewEvent(sesamy.EventNameTutorialComplete, p))
}
