package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type ViewItem sesamy2.Event[params.ViewItem[params.Item]]

func NewViewItem(p params.ViewItem[params.Item]) ViewItem {
	return ViewItem(sesamy2.NewEvent(sesamy2.EventNameViewItem, p))
}
