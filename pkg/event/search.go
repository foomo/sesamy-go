package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type Search sesamy2.Event[params.Search]

func NewSearch(p params.Search) Search {
	return Search(sesamy2.NewEvent(sesamy2.EventNameSearch, p))
}
