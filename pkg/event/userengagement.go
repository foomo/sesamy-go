package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// UserEngagement is the GA4 user_engagement event.
type UserEngagement sesamy.Event[params.UserEngagement]

// NewUserEngagement returns an user_engagement event with the parameters p.
func NewUserEngagement(p params.UserEngagement) sesamy.Event[params.UserEngagement] {
	return sesamy.NewEvent(sesamy.EventNameUserEngagement, p)
}
