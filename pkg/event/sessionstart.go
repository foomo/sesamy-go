package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type SessionStart sesamy.Event[params.SessionStart]

func NewSessionStart(p params.SessionStart) sesamy.Event[params.SessionStart] {
	return sesamy.NewEvent(sesamy.EventNameSessionStart, p)
}
