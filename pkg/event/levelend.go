package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type LevelEnd sesamy.Event[params.LevelEnd]

func NewLevelEnd(p params.LevelEnd) LevelEnd {
	return LevelEnd(sesamy.NewEvent(sesamy.EventNameLevelEnd, p))
}
