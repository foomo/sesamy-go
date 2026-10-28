package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// UnlockAchievement is the GA4 unlock_achievement event.
type UnlockAchievement sesamy.Event[params.UnlockAchievement]

// NewUnlockAchievement returns an unlock_achievement event with the parameters p.
func NewUnlockAchievement(p params.UnlockAchievement) sesamy.Event[params.UnlockAchievement] {
	return sesamy.NewEvent(sesamy.EventNameUnlockAchievement, p)
}
