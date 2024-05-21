package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type Search sesamy.Event[params.Search]

func NewSearch(p params.Search) Search {
	return Search(sesamy.NewEvent(sesamy.EventNameSearch, p))
}
