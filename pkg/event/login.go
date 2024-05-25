package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type Login sesamy.Event[params.Login]

func NewLogin(p params.Login) sesamy.Event[params.Login] {
	return sesamy.NewEvent(sesamy.EventNameLogin, p)
}
