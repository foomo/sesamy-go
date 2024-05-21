package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type ViewSearchResults sesamy2.Event[params.ViewSearchResults]

func NewViewSearchResults(p params.ViewSearchResults) ViewSearchResults {
	return ViewSearchResults(sesamy2.NewEvent(sesamy2.EventNameViewSearchResults, p))
}
