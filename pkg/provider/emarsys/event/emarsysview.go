package event

import (
	sesamyparams "github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/provider/emarsys/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

const EventNameEmarsysView sesamy.EventName = "emarsys_view"

type EmarsysView sesamy.Event[params.EmarsysView[sesamyparams.Item]]

func NewEmarsysView(p params.EmarsysView[sesamyparams.Item]) sesamy.Event[params.EmarsysView[sesamyparams.Item]] {
	return sesamy.NewEvent(EventNameEmarsysView, p)
}
