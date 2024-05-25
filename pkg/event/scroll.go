package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type Scroll sesamy.Event[params.Scroll]

func NewScroll(p params.Scroll) sesamy.Event[params.Scroll] {
	return sesamy.NewEvent(sesamy.EventNameScroll, p)
}
