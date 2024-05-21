package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type VideoProgress sesamy.Event[params.VideoProgress]

func NewVideoProgress(p params.VideoProgress) VideoProgress {
	return VideoProgress(sesamy.NewEvent(sesamy.EventNameVideoProgress, p))
}
