package params

// SelectPromotion https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference/events#select_promotion
type SelectPromotion[Item any] struct {
	CreativeName  string `json:"creative_name,omitempty"`
	CreativeSlot  string `json:"creative_slot,omitempty"`
	PromotionID   string `json:"item_list_id,omitempty"`
	PromotionName string `json:"item_list_name,omitempty"`
	Items         []Item `json:"items,omitempty"`
}
