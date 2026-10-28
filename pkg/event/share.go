package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// Share is the GA4 share event.
type Share sesamy.Event[params.Share]

// NewShare returns a share event with the parameters p.
func NewShare(p params.Share) sesamy.Event[params.Share] {
	return sesamy.NewEvent(sesamy.EventNameShare, p)
}
