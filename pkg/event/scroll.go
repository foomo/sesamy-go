package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type Scroll sesamy2.Event[params.Scroll]

func NewScroll(p params.Scroll) Scroll {
	return Scroll(sesamy2.NewEvent(sesamy2.EventNameScroll, p))
}
