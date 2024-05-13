package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/encoding/mpv2/params"
)

type SelectItem sesamy.Event[params.SelectItem[params.Item]]

func NewSelectItem(p params.SelectItem[params.Item]) SelectItem {
	return SelectItem(sesamy.NewEvent(sesamy.EventNameSelectItem, p))
}
