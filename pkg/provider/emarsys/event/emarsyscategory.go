package event

import (
	"github.com/foomo/sesamy-go/pkg/provider/emarsys/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// EventNameEmarsysCategory is the name of the emarsys_category event.
const EventNameEmarsysCategory sesamy.EventName = "emarsys_category"

// EmarsysCategory is the emarsys_category event.
type EmarsysCategory sesamy.Event[params.EmarsysCategory]

// NewEmarsysCategory returns an emarsys_category event with the parameters p.
func NewEmarsysCategory(p params.EmarsysCategory) sesamy.Event[params.EmarsysCategory] {
	return sesamy.NewEvent(EventNameEmarsysCategory, p)
}
