package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type Scroll sesamy.Event[params.Scroll]

func NewScroll(p params.Scroll) Scroll {
	return Scroll(sesamy.NewEvent(sesamy.EventNameScroll, p))
}
