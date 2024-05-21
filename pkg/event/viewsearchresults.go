package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type ViewSearchResults sesamy.Event[params.ViewSearchResults]

func NewViewSearchResults(p params.ViewSearchResults) ViewSearchResults {
	return ViewSearchResults(sesamy.NewEvent(sesamy.EventNameViewSearchResults, p))
}
