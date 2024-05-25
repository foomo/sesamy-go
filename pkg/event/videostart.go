package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type VideoStart sesamy.Event[params.VideoStart]

func NewVideoStart(p params.VideoStart) sesamy.Event[params.VideoStart] {
	return sesamy.NewEvent(sesamy.EventNameVideoStart, p)
}
