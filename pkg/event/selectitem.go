package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// SelectItem is the GA4 select_item event.
type SelectItem sesamy.Event[params.SelectItem[params.Item]]

// NewSelectItem returns a select_item event with the parameters p.
func NewSelectItem(p params.SelectItem[params.Item]) sesamy.Event[params.SelectItem[params.Item]] {
	return sesamy.NewEvent(sesamy.EventNameSelectItem, p)
}
