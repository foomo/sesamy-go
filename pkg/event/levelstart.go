package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// LevelStart is the GA4 level_start event.
type LevelStart sesamy.Event[params.LevelStart]

// NewLevelStart returns a level_start event with the parameters p.
func NewLevelStart(p params.LevelStart) sesamy.Event[params.LevelStart] {
	return sesamy.NewEvent(sesamy.EventNameLevelStart, p)
}
