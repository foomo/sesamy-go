package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// ViewItem is the GA4 view_item event.
type ViewItem sesamy.Event[params.ViewItem[params.Item]]

// NewViewItem returns a view_item event with the parameters p.
func NewViewItem(p params.ViewItem[params.Item]) sesamy.Event[params.ViewItem[params.Item]] {
	return sesamy.NewEvent(sesamy.EventNameViewItem, p)
}
