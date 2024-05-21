package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type Share sesamy.Event[params.Share]

func NewShare(p params.Share) Share {
	return Share(sesamy.NewEvent(sesamy.EventNameShare, p))
}
