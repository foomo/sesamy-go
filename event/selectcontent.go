package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/event/params"
)

type SelectContent sesamy.Event[params.SelectContent]

func NewSelectContent(p params.SelectContent) SelectContent {
	return SelectContent(sesamy.NewEvent(sesamy.EventNameSelectContent, p))
}
