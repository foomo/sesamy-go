package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/encoding/mpv2/params"
)

type ViewSearchResults sesamy.Event[params.ViewSearchResults]

func NewViewSearchResults(p params.ViewSearchResults) ViewSearchResults {
	return ViewSearchResults(sesamy.NewEvent(sesamy.EventNameViewSearchResults, p))
}
