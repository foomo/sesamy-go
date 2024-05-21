package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type VideoStart sesamy2.Event[params.VideoStart]

func NewVideoStart(p params.VideoStart) VideoStart {
	return VideoStart(sesamy2.NewEvent(sesamy2.EventNameVideoStart, p))
}
