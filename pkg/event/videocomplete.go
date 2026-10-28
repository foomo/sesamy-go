package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// VideoComplete is the GA4 video_complete event.
type VideoComplete sesamy.Event[params.VideoComplete]

// NewVideoComplete returns a video_complete event with the parameters p.
func NewVideoComplete(p params.VideoComplete) sesamy.Event[params.VideoComplete] {
	return sesamy.NewEvent(sesamy.EventNameVideoComplete, p)
}
