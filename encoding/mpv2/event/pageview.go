package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/encoding/mpv2/params"
)

type PageView sesamy.Event[params.PageView]

func NewPageView(p params.PageView) PageView {
	return PageView(sesamy.NewEvent(sesamy.EventNamePageView, p))
}
