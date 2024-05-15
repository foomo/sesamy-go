package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/event/params"
)

type Scroll sesamy.Event[params.Scroll]

func NewScroll(p params.Scroll) Scroll {
	return Scroll(sesamy.NewEvent(sesamy.EventNameScroll, p))
}
