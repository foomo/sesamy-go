package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// JoinGroup is the GA4 join_group event.
type JoinGroup sesamy.Event[params.JoinGroup]

// NewJoinGroup returns a join_group event with the parameters p.
func NewJoinGroup(p params.JoinGroup) sesamy.Event[params.JoinGroup] {
	return sesamy.NewEvent(sesamy.EventNameJoinGroup, p)
}
