package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type UserEngagement sesamy2.Event[params.UserEngagement]

func NewUserEngagement(p params.UserEngagement) UserEngagement {
	return UserEngagement(sesamy2.NewEvent(sesamy2.EventNameUserEngagement, p))
}
