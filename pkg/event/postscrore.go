package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type PostScore sesamy.Event[params.PostScore]

func NewPostScore(p params.PostScore) PostScore {
	return PostScore(sesamy.NewEvent(sesamy.EventNamePostScore, p))
}
