package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/event/params"
)

type Search sesamy.Event[params.Search]

func NewSearch(p params.Search) Search {
	return Search(sesamy.NewEvent(sesamy.EventNameSearch, p))
}
