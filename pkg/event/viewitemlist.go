package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type ViewItemList sesamy2.Event[params.ViewItemList[params.Item]]

func NewViewItemList(p params.ViewItemList[params.Item]) ViewItemList {
	return ViewItemList(sesamy2.NewEvent(sesamy2.EventNameViewItemList, p))
}
