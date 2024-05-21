package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type JoinGroup sesamy.Event[params.JoinGroup]

func NewJoinGroup(p params.JoinGroup) JoinGroup {
	return JoinGroup(sesamy.NewEvent(sesamy.EventNameJoinGroup, p))
}
