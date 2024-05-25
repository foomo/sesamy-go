package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type JoinGroup sesamy.Event[params.JoinGroup]

func NewJoinGroup(p params.JoinGroup) sesamy.Event[params.JoinGroup] {
	return sesamy.NewEvent(sesamy.EventNameJoinGroup, p)
}
