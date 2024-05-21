package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type VideoComplete sesamy.Event[params.VideoComplete]

func NewVideoComplete(p params.VideoComplete) VideoComplete {
	return VideoComplete(sesamy.NewEvent(sesamy.EventNameVideoComplete, p))
}
