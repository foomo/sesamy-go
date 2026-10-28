package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// LevelEnd is the GA4 level_end event.
type LevelEnd sesamy.Event[params.LevelEnd]

// NewLevelEnd returns a level_end event with the parameters p.
func NewLevelEnd(p params.LevelEnd) sesamy.Event[params.LevelEnd] {
	return sesamy.NewEvent(sesamy.EventNameLevelEnd, p)
}
