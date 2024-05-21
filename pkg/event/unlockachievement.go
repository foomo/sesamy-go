package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type UnlockAchievement sesamy.Event[params.UnlockAchievement]

func NewUnlockArchievement(p params.UnlockAchievement) UnlockAchievement {
	return UnlockAchievement(sesamy.NewEvent(sesamy.EventNameUnlockAchievement, p))
}
