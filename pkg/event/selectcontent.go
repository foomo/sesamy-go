package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// SelectContent is the GA4 select_content event.
type SelectContent sesamy.Event[params.SelectContent]

// NewSelectContent returns a select_content event with the parameters p.
func NewSelectContent(p params.SelectContent) sesamy.Event[params.SelectContent] {
	return sesamy.NewEvent(sesamy.EventNameSelectContent, p)
}
