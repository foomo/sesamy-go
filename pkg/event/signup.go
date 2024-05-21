package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type SignUp sesamy.Event[params.SignUp]

func NewSignUp(p params.SignUp) SignUp {
	return SignUp(sesamy.NewEvent(sesamy.EventNameSignUp, p))
}
