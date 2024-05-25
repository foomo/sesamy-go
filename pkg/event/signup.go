package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type SignUp sesamy.Event[params.SignUp]

func NewSignUp(p params.SignUp) sesamy.Event[params.SignUp] {
	return sesamy.NewEvent(sesamy.EventNameSignUp, p)
}
