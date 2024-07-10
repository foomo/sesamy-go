package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type PageView sesamy.Event[params.PageView]

func NewPageView(p params.PageView) sesamy.Event[params.PageView] {
	return sesamy.NewEvent(sesamy.EventNamePageView, p)
}
