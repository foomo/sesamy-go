package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type LevelEnd sesamy.Event[params.LevelEnd]

func NewLevelEnd(p params.LevelEnd) sesamy.Event[params.LevelEnd] {
	return sesamy.NewEvent(sesamy.EventNameLevelEnd, p)
}
