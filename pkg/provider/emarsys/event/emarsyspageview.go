package event

import (
	sesamyparams "github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/provider/emarsys/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

const EventNameEmarsysPageView sesamy.EventName = "emarsys_page_view"

type EmarsysPageView sesamy.Event[params.EmarsysPageView[sesamyparams.Item]]

func NewEmarsysPageView(p params.EmarsysPageView[sesamyparams.Item]) sesamy.Event[params.EmarsysPageView[sesamyparams.Item]] {
	return sesamy.NewEvent(EventNameEmarsysPageView, p)
}
