package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// Login is the GA4 login event.
type Login sesamy.Event[params.Login]

// NewLogin returns a login event with the parameters p.
func NewLogin(p params.Login) sesamy.Event[params.Login] {
	return sesamy.NewEvent(sesamy.EventNameLogin, p)
}
