package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type LevelEnd sesamy2.Event[params.LevelEnd]

func NewLevelEnd(p params.LevelEnd) LevelEnd {
	return LevelEnd(sesamy2.NewEvent(sesamy2.EventNameLevelEnd, p))
}
