package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// PageView is the GA4 page_view event.
type PageView sesamy.Event[params.PageView]

// NewPageView returns a page_view event with the parameters p.
func NewPageView(p params.PageView) sesamy.Event[params.PageView] {
	return sesamy.NewEvent(sesamy.EventNamePageView, p)
}
