package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// VideoProgress is the GA4 video_progress event.
type VideoProgress sesamy.Event[params.VideoProgress]

// NewVideoProgress returns a video_progress event with the parameters p.
func NewVideoProgress(p params.VideoProgress) sesamy.Event[params.VideoProgress] {
	return sesamy.NewEvent(sesamy.EventNameVideoProgress, p)
}
