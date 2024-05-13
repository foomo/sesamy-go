package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/encoding/mpv2/params"
)

type Share sesamy.Event[params.Share]

func NewShare(p params.Share) Share {
	return Share(sesamy.NewEvent(sesamy.EventNameShare, p))
}
