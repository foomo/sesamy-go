package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// PostScore is the GA4 post_score event.
type PostScore sesamy.Event[params.PostScore]

// NewPostScore returns a post_score event with the parameters p.
func NewPostScore(p params.PostScore) sesamy.Event[params.PostScore] {
	return sesamy.NewEvent(sesamy.EventNamePostScore, p)
}
