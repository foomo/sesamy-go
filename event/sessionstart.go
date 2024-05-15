package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/event/params"
)

type SessionStart sesamy.Event[params.SessionStart]

func NewSessionStart(p params.SessionStart) SessionStart {
	return SessionStart(sesamy.NewEvent(sesamy.EventNameSessionStart, p))
}
