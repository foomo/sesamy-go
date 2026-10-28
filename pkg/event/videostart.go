package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// VideoStart is the GA4 video_start event.
type VideoStart sesamy.Event[params.VideoStart]

// NewVideoStart returns a video_start event with the parameters p.
func NewVideoStart(p params.VideoStart) sesamy.Event[params.VideoStart] {
	return sesamy.NewEvent(sesamy.EventNameVideoStart, p)
}
