package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type LevelUp sesamy2.Event[params.LevelUp]

func NewLevelUp(p params.LevelUp) LevelUp {
	return LevelUp(sesamy2.NewEvent(sesamy2.EventNameLevelUp, p))
}
