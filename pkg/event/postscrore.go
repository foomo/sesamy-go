package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type PostScore sesamy.Event[params.PostScore]

func NewPostScore(p params.PostScore) sesamy.Event[params.PostScore] {
	return sesamy.NewEvent(sesamy.EventNamePostScore, p)
}
