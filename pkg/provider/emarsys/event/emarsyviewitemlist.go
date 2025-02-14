package event

import (
	"github.com/foomo/sesamy-go/pkg/provider/emarsys/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

const EventNameEmarsysViewItemList sesamy.EventName = "emarsys_view_item_list"

type EmarsysViewItemList sesamy.Event[params.EmarsysViewItemList]

func NewEmarsysViewItemList(p params.EmarsysViewItemList) sesamy.Event[params.EmarsysViewItemList] {
	return sesamy.NewEvent(EventNameEmarsysViewItemList, p)
}
