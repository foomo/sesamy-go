package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type UserEngagement sesamy.Event[params.UserEngagement]

func NewUserEngagement(p params.UserEngagement) sesamy.Event[params.UserEngagement] {
	return sesamy.NewEvent(sesamy.EventNameUserEngagement, p)
}
