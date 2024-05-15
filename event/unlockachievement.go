package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/event/params"
)

type UnlockAchievement sesamy.Event[params.UnlockAchievement]

func NewUnlockArchievement(p params.UnlockAchievement) UnlockAchievement {
	return UnlockAchievement(sesamy.NewEvent(sesamy.EventNameUnlockArchievement, p))
}
