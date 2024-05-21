package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type Login sesamy2.Event[params.Login]

func NewLogin(p params.Login) Login {
	return Login(sesamy2.NewEvent(sesamy2.EventNameLogin, p))
}
