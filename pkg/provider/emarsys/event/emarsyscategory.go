package event

import (
	"github.com/foomo/sesamy-go/pkg/provider/emarsys/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

const EventNameEmarsysCategory sesamy.EventName = "emarsys_category"

type EmarsysCategory sesamy.Event[params.EmarsysCategory]

func NewEmarsysCategory(p params.EmarsysCategory) sesamy.Event[params.EmarsysCategory] {
	return sesamy.NewEvent(EventNameEmarsysCategory, p)
}
