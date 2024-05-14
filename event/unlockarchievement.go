package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/event/params"
)

type UnlockArchievement sesamy.Event[params.UnlockArchievement]

func NewUnlockArchievement(p params.UnlockArchievement) UnlockArchievement {
	return UnlockArchievement(sesamy.NewEvent(sesamy.EventNameUnlockArchievement, p))
}
