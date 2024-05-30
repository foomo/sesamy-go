package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type UnlockAchievement sesamy.Event[params.UnlockAchievement]

func NewUnlockAchievement(p params.UnlockAchievement) sesamy.Event[params.UnlockAchievement] {
	return sesamy.NewEvent(sesamy.EventNameUnlockAchievement, p)
}
