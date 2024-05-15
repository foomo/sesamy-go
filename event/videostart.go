package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/event/params"
)

type VideoStart sesamy.Event[params.VideoStart]

func NewVideoStart(p params.VideoStart) VideoStart {
	return VideoStart(sesamy.NewEvent(sesamy.EventNameVideoStart, p))
}
