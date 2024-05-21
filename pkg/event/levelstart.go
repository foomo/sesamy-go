package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type LevelStart sesamy2.Event[params.LevelStart]

func NewLevelStart(p params.LevelStart) LevelStart {
	return LevelStart(sesamy2.NewEvent(sesamy2.EventNameLevelStart, p))
}
