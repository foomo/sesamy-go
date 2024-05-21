package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type Share sesamy2.Event[params.Share]

func NewShare(p params.Share) Share {
	return Share(sesamy2.NewEvent(sesamy2.EventNameShare, p))
}
