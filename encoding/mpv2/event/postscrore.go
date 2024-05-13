package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/encoding/mpv2/params"
)

type PostScore sesamy.Event[params.PostScore]

func NewPostScore(p params.PostScore) PostScore {
	return PostScore(sesamy.NewEvent(sesamy.EventNamePostScore, p))
}
