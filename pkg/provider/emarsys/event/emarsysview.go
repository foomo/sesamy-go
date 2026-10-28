package event

import (
	sesamyparams "github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/provider/emarsys/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// EventNameEmarsysView is the name of the emarsys_view event.
const EventNameEmarsysView sesamy.EventName = "emarsys_view"

// EmarsysView is the emarsys_view event.
type EmarsysView sesamy.Event[params.EmarsysView[sesamyparams.Item]]

// NewEmarsysView returns an emarsys_view event with the parameters p.
func NewEmarsysView(p params.EmarsysView[sesamyparams.Item]) sesamy.Event[params.EmarsysView[sesamyparams.Item]] {
	return sesamy.NewEvent(EventNameEmarsysView, p)
}
