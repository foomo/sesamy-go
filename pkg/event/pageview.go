package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type PageView sesamy2.Event[params.PageView]

func NewPageView(p params.PageView) PageView {
	return PageView(sesamy2.NewEvent(sesamy2.EventNamePageView, p))
}
