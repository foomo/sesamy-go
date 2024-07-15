package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type Exception sesamy.Event[params.Exception]

func NewException(p params.Exception) sesamy.Event[params.Exception] {
	return sesamy.NewEvent(sesamy.EventNameException, p)
}
