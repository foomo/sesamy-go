package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/encoding/mpv2/params"
)

type ViewItemList sesamy.Event[params.ViewItemList[params.Item]]

func NewViewItemList(p params.ViewItemList[params.Item]) ViewItemList {
	return ViewItemList(sesamy.NewEvent(sesamy.EventNameViewItemList, p))
}
