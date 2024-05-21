package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type SelectItem sesamy.Event[params.SelectItem[params.Item]]

func NewSelectItem(p params.SelectItem[params.Item]) SelectItem {
	return SelectItem(sesamy.NewEvent(sesamy.EventNameSelectItem, p))
}
