package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type UnlockAchievement sesamy2.Event[params.UnlockAchievement]

func NewUnlockArchievement(p params.UnlockAchievement) UnlockAchievement {
	return UnlockAchievement(sesamy2.NewEvent(sesamy2.EventNameUnlockAchievement, p))
}
