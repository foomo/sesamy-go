package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type ViewItem sesamy.Event[params.ViewItem[params.Item]]

func NewViewItem(p params.ViewItem[params.Item]) ViewItem {
	return ViewItem(sesamy.NewEvent(sesamy.EventNameViewItem, p))
}
