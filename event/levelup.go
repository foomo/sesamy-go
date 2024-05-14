package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/event/params"
)

type LevelUp sesamy.Event[params.LevelUp]

func NewLevelUp(p params.LevelUp) LevelUp {
	return LevelUp(sesamy.NewEvent(sesamy.EventNameLevelUp, p))
}
