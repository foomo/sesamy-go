package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type LevelStart sesamy.Event[params.LevelStart]

func NewLevelStart(p params.LevelStart) LevelStart {
	return LevelStart(sesamy.NewEvent(sesamy.EventNameLevelStart, p))
}
