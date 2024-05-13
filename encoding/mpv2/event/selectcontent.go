package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/encoding/mpv2/params"
)

type SelectContent sesamy.Event[params.SelectContent]

func NewSelectContent(p params.SelectContent) SelectContent {
	return SelectContent(sesamy.NewEvent(sesamy.EventNameSelectContent, p))
}
