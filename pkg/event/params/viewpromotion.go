package params

// ViewPromotion https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference/events#view_promotion
type ViewPromotion[I any] struct {
	CreativeName  string `json:"creative_name,omitempty"`
	CreativeSlot  string `json:"creative_slot,omitempty"`
	PromotionID   string `json:"promotion_id,omitempty"`
	PromotionName string `json:"promotion_name,omitempty"`
	Items         []I    `json:"items,omitempty"`
}
