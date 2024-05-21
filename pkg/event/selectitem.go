package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type SelectItem sesamy2.Event[params.SelectItem[params.Item]]

func NewSelectItem(p params.SelectItem[params.Item]) SelectItem {
	return SelectItem(sesamy2.NewEvent(sesamy2.EventNameSelectItem, p))
}
