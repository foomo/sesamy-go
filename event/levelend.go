package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/event/params"
)

type LevelEnd sesamy.Event[params.LevelEnd]

func NewLevelEnd(p params.LevelEnd) LevelEnd {
	return LevelEnd(sesamy.NewEvent(sesamy.EventNameLevelEnd, p))
}
