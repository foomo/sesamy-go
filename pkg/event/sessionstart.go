package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type SessionStart sesamy.Event[params.SessionStart]

func NewSessionStart(p params.SessionStart) SessionStart {
	return SessionStart(sesamy.NewEvent(sesamy.EventNameSessionStart, p))
}
