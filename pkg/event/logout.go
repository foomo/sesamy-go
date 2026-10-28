package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// Logout is the GA4 logout event.
type Logout sesamy.Event[params.Logout]

// NewLogout returns a logout event with the parameters p.
func NewLogout(p params.Logout) sesamy.Event[params.Logout] {
	return sesamy.NewEvent(sesamy.EventNameLogout, p)
}
