package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/event/params"
)

type PageView sesamy.Event[params.PageView]

func NewPageView(p params.PageView) PageView {
	return PageView(sesamy.NewEvent(sesamy.EventNamePageView, p))
}
