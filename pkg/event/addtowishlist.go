package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// AddToWishlist is the GA4 add_to_wishlist event.
type AddToWishlist sesamy.Event[params.AddToWishlist[params.Item]]

// NewAddToWishlist returns an add_to_wishlist event with the parameters p.
func NewAddToWishlist(p params.AddToWishlist[params.Item]) sesamy.Event[params.AddToWishlist[params.Item]] {
	return sesamy.NewEvent(sesamy.EventNameAddToWishlist, p)
}
