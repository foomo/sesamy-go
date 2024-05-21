package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type Login sesamy.Event[params.Login]

func NewLogin(p params.Login) Login {
	return Login(sesamy.NewEvent(sesamy.EventNameLogin, p))
}
