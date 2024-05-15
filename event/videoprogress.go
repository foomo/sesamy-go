package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/event/params"
)

type VideoProgress sesamy.Event[params.VideoProgress]

func NewVideoProgress(p params.VideoProgress) VideoProgress {
	return VideoProgress(sesamy.NewEvent(sesamy.EventNameVideoProgress, p))
}
