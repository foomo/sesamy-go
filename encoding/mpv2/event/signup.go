package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/encoding/mpv2/params"
)

type SignUp sesamy.Event[params.SignUp]

func NewSignUp(p params.SignUp) SignUp {
	return SignUp(sesamy.NewEvent(sesamy.EventNameSignUp, p))
}
