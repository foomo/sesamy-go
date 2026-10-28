package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// TutorialComplete is the GA4 tutorial_complete event.
type TutorialComplete sesamy.Event[params.TutorialComplete]

// NewTutorialComplete returns a tutorial_complete event with the parameters p.
func NewTutorialComplete(p params.TutorialComplete) sesamy.Event[params.TutorialComplete] {
	return sesamy.NewEvent(sesamy.EventNameTutorialComplete, p)
}
