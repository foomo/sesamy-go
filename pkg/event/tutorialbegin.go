package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// TutorialBegin is the GA4 tutorial_begin event.
type TutorialBegin sesamy.Event[params.TutorialBegin]

// NewTutorialBegin returns a tutorial_begin event with the parameters p.
func NewTutorialBegin(p params.TutorialBegin) sesamy.Event[params.TutorialBegin] {
	return sesamy.NewEvent(sesamy.EventNameTutorialBegin, p)
}
