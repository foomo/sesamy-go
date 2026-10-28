package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// ViewItemList is the GA4 view_item_list event.
type ViewItemList sesamy.Event[params.ViewItemList[params.Item]]

// NewViewItemList returns a view_item_list event with the parameters p.
func NewViewItemList(p params.ViewItemList[params.Item]) sesamy.Event[params.ViewItemList[params.Item]] {
	return sesamy.NewEvent(sesamy.EventNameViewItemList, p)
}
