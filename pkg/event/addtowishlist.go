package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type AddToWishlist sesamy2.Event[params.AddToWishlist[params.Item]]

func NewAddToWishlist(p params.AddToWishlist[params.Item]) AddToWishlist {
	return AddToWishlist(sesamy2.NewEvent(sesamy2.EventNameAddToWishlist, p))
}
