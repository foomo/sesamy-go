package event

import (
	sesamyparams "github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/provider/emarsys/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

const EventNameEmarsysViewItem sesamy.EventName = "emarsys_view_item"

type EmarsysViewItem sesamy.Event[params.EmarsysViewItem[sesamyparams.Item]]

func NewEmarsysViewItem(p params.EmarsysViewItem[sesamyparams.Item]) sesamy.Event[params.EmarsysViewItem[sesamyparams.Item]] {
	return sesamy.NewEvent(EventNameEmarsysViewItem, p)
}
