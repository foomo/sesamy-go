package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type Logout sesamy.Event[params.Logout]

func NewLogout(p params.Logout) sesamy.Event[params.Logout] {
	return sesamy.NewEvent(sesamy.EventNameLogout, p)
}
