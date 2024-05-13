package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/encoding/mpv2/params"
)

type ViewItem sesamy.Event[params.ViewItem[params.Item]]

func NewViewItem(p params.ViewItem[params.Item]) ViewItem {
	return ViewItem(sesamy.NewEvent(sesamy.EventNameViewItem, p))
}
