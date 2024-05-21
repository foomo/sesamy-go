package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type AddToWishlist sesamy.Event[params.AddToWishlist[params.Item]]

func NewAddToWishlist(p params.AddToWishlist[params.Item]) AddToWishlist {
	return AddToWishlist(sesamy.NewEvent(sesamy.EventNameAddToWishlist, p))
}
