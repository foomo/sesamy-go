package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type SelectContent sesamy2.Event[params.SelectContent]

func NewSelectContent(p params.SelectContent) SelectContent {
	return SelectContent(sesamy2.NewEvent(sesamy2.EventNameSelectContent, p))
}
