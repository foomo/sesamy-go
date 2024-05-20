package params

type Item struct {
	Affiliation   string  `json:"affiliation,omitempty"`
	Coupon        string  `json:"coupon,omitempty"`
	CreativeName  string  `json:"creative_name,omitempty"`
	CreativeSlot  string  `json:"creative_slot,omitempty"`
	Discount      float64 `json:"discount,omitempty"`
	Index         int     `json:"index,omitempty"`
	ItemBrand     string  `json:"item_brand,omitempty"`
	ItemCategory  string  `json:"item_category,omitempty"`
	ItemCategory2 string  `json:"item_category2,omitempty"`
	ItemCategory3 string  `json:"item_category3,omitempty"`
	ItemCategory4 string  `json:"item_category4,omitempty"`
	ItemCategory5 string  `json:"item_category5,omitempty"`
	ItemID        string  `json:"item_id,omitempty"`
	ItemListID    string  `json:"item_list_id,omitempty"`
	ItemListName  string  `json:"item_list_name,omitempty"`
	ItemName      string  `json:"item_name,omitempty"`
	ItemVariant   string  `json:"item_variant,omitempty"`
	LocationID    string  `json:"location_id,omitempty"`
	Price         float64 `json:"price,omitempty"`
	PromotionID   string  `json:"promotion_id,omitempty"`
	PromotionName string  `json:"promotion_name,omitempty"`
	Quantity      float64 `json:"quantity,omitempty"`
}
