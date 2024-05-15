package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/event/params"
)

type VideoComplete sesamy.Event[params.VideoComplete]

func NewVideoComplete(p params.VideoComplete) VideoComplete {
	return VideoComplete(sesamy.NewEvent(sesamy.EventNameVideoComplete, p))
}
