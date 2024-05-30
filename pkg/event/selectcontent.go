package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type SelectContent sesamy.Event[params.SelectContent]

func NewSelectContent(p params.SelectContent) sesamy.Event[params.SelectContent] {
	return sesamy.NewEvent(sesamy.EventNameSelectContent, p)
}
