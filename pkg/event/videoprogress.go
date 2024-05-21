package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type VideoProgress sesamy2.Event[params.VideoProgress]

func NewVideoProgress(p params.VideoProgress) VideoProgress {
	return VideoProgress(sesamy2.NewEvent(sesamy2.EventNameVideoProgress, p))
}
