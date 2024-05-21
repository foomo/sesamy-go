package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type PageView sesamy.Event[params.PageView]

func NewPageView(p params.PageView) PageView {
	return PageView(sesamy.NewEvent(sesamy.EventNamePageView, p))
}
