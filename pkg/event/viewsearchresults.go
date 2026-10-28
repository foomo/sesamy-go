package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// ViewSearchResults is the GA4 view_search_results event.
type ViewSearchResults sesamy.Event[params.ViewSearchResults]

// NewViewSearchResults returns a view_search_results event with the parameters p.
func NewViewSearchResults(p params.ViewSearchResults) sesamy.Event[params.ViewSearchResults] {
	return sesamy.NewEvent(sesamy.EventNameViewSearchResults, p)
}
