package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/event/params"
)

type ViewItemList sesamy.Event[params.ViewItemList[params.Item]]

func NewViewItemList(p params.ViewItemList[params.Item]) ViewItemList {
	return ViewItemList(sesamy.NewEvent(sesamy.EventNameViewItemList, p))
}
