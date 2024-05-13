package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/encoding/mpv2/params"
)

type JoinGroup sesamy.Event[params.JoinGroup]

func NewJoinGroup(p params.JoinGroup) JoinGroup {
	return JoinGroup(sesamy.NewEvent(sesamy.EventNameJoinGroup, p))
}
