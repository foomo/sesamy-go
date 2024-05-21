package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type VideoStart sesamy.Event[params.VideoStart]

func NewVideoStart(p params.VideoStart) VideoStart {
	return VideoStart(sesamy.NewEvent(sesamy.EventNameVideoStart, p))
}
