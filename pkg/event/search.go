package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// Search is the GA4 search event.
type Search sesamy.Event[params.Search]

// NewSearch returns a search event with the parameters p.
func NewSearch(p params.Search) sesamy.Event[params.Search] {
	return sesamy.NewEvent(sesamy.EventNameSearch, p)
}
