package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type VideoComplete sesamy2.Event[params.VideoComplete]

func NewVideoComplete(p params.VideoComplete) VideoComplete {
	return VideoComplete(sesamy2.NewEvent(sesamy2.EventNameVideoComplete, p))
}
