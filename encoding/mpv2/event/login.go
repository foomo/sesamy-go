package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/encoding/mpv2/params"
)

type Login sesamy.Event[params.Login]

func NewLogin(p params.Login) Login {
	return Login(sesamy.NewEvent(sesamy.EventNameLogin, p))
}
