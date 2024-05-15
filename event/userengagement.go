package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/event/params"
)

type UserEngagement sesamy.Event[params.UserEngagement]

func NewUserEngagement(p params.UserEngagement) UserEngagement {
	return UserEngagement(sesamy.NewEvent(sesamy.EventNameUserEngagement, p))
}
