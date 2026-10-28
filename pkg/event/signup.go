package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// SignUp is the GA4 sign_up event.
type SignUp sesamy.Event[params.SignUp]

// NewSignUp returns a sign_up event with the parameters p.
func NewSignUp(p params.SignUp) sesamy.Event[params.SignUp] {
	return sesamy.NewEvent(sesamy.EventNameSignUp, p)
}
