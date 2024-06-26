package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type VideoComplete sesamy.Event[params.VideoComplete]

func NewVideoComplete(p params.VideoComplete) sesamy.Event[params.VideoComplete] {
	return sesamy.NewEvent(sesamy.EventNameVideoComplete, p)
}
