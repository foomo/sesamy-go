package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type JoinGroup sesamy2.Event[params.JoinGroup]

func NewJoinGroup(p params.JoinGroup) JoinGroup {
	return JoinGroup(sesamy2.NewEvent(sesamy2.EventNameJoinGroup, p))
}
