package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/encoding/mpv2/params"
)

type LevelStart sesamy.Event[params.LevelStart]

func NewLevelStart(p params.LevelStart) LevelStart {
	return LevelStart(sesamy.NewEvent(sesamy.EventNameLevelStart, p))
}
