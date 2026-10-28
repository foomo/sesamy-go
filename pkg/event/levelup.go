package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// LevelUp is the GA4 level_up event.
type LevelUp sesamy.Event[params.LevelUp]

// NewLevelUp returns a level_up event with the parameters p.
func NewLevelUp(p params.LevelUp) sesamy.Event[params.LevelUp] {
	return sesamy.NewEvent(sesamy.EventNameLevelUp, p)
}
