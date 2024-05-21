package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type LevelUp sesamy.Event[params.LevelUp]

func NewLevelUp(p params.LevelUp) LevelUp {
	return LevelUp(sesamy.NewEvent(sesamy.EventNameLevelUp, p))
}
