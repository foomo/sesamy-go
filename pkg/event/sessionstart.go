package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type SessionStart sesamy2.Event[params.SessionStart]

func NewSessionStart(p params.SessionStart) SessionStart {
	return SessionStart(sesamy2.NewEvent(sesamy2.EventNameSessionStart, p))
}
