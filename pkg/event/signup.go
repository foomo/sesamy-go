package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type SignUp sesamy2.Event[params.SignUp]

func NewSignUp(p params.SignUp) SignUp {
	return SignUp(sesamy2.NewEvent(sesamy2.EventNameSignUp, p))
}
