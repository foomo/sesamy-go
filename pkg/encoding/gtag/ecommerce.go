package gtag

import (
	"github.com/foomo/gostandards/iso4217"
)

type ECommerce struct {
	// Currency Code. ISO 4217
	// Example: JPY
	Currency *iso4217.Currency `json:"currency,omitempty" gtag:"cu,omitempty"`
	// Example:
	Items []*Item `json:"items,omitempty" gtag:"pr,omitempty"`
	// Promotion Impression/Click Tracking. Promotion Id
	// Example: summer-offer
	PromotionID *string `json:"promotion_id,omitempty" gtag:"pi,omitempty"`
	// Promotion Impression/Click Tracking. Promotion Name
	// Example: summer-offer
	PromotionName *string `json:"promotion_name,omitempty" gtag:"pn,omitempty"`
	// Promotion Impression/Click Tracking. Creative Name
	// Example: red-car
	// CreativeName *string `json:"//,omitempty" gtag:"cn,omitempty"`
	// Promotion Impression/Click Tracking. Promotion Slot / Position
	// Example: slide-3
	// CreativeSlot *string `json:"//,omitempty" gtag:"cs,omitempty"`
	// Google Place ID: Refer to: https://developers.google.com/maps/documentation/places/web-service/place-id . Seems to be inherited from Firebase, not sure about the current use on GA4
	// Example: ChIJiyj437sx3YAR9kUWC8QkLzQ
	LocationID *string `json:"location_id,omitempty" gtag:"lo,omitempty"`
	// If the current event is set as a conversion on the admin interacted the evfent will have this value present
	// Example: 1
	IsConversion *string `json:"is_conversion,omitempty" gtag:"_c,omitempty"`
}
