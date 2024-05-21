package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type PostScore sesamy2.Event[params.PostScore]

func NewPostScore(p params.PostScore) PostScore {
	return PostScore(sesamy2.NewEvent(sesamy2.EventNamePostScore, p))
}
