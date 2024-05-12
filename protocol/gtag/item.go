package gtag

type Item struct {
	// Example: 12345
	ItemID *string `json:"item_id,omitempty" gtag:"id,omitempty"`
	// Example: Stan and Friends Tee
	ItemName *string `json:"item_name,omitempty" gtag:"nm,omitempty"`
	// Example: Google
	ItemBrand *string `json:"item_brand,omitempty" gtag:"br,omitempty"`
	// Example: men
	ItemCategory *string `json:"item_category,omitempty" gtag:"ca,omitempty"`
	// Example: t-shirts
	ItemCategory2 *string `json:"item_category2,omitempty" gtag:"c2,omitempty"`
	// Example: men
	ItemCategory3 *string `json:"item_category3,omitempty" gtag:"c3,omitempty"`
	// Example: men
	ItemCategory4 *string `json:"item_category4,omitempty" gtag:"c4,omitempty"`
	// Example: men
	ItemCategory5 *string `json:"item_category5,omitempty" gtag:"c5,omitempty"`
	// Example: Yellow
	ItemVariant *string `json:"item_variant,omitempty" gtag:"va,omitempty"`
	// Example: 123.45
	Price *string `json:"price,omitempty" gtag:"pr,omitempty"`
	// Example: 1
	Quantity *string `json:"quantity,omitempty" gtag:"qt,omitempty"`
	// Example: 50%OFF
	Coupon *string `json:"coupon,omitempty" gtag:"cp,omitempty"`
	// Example: cross-selling: mens
	ItemListName *string `json:"item_list_name,omitempty" gtag:"ln,omitempty"`
	// Example: 10
	ItemListPosition *string `json:"item_list_position,omitempty" gtag:"lp,omitempty"`
	// Example: id-mens-123
	ItemListID *string `json:"item_list_id,omitempty" gtag:"li,omitempty"`
	// Example: 10.00
	Discount *string `json:"discount,omitempty" gtag:"ds,omitempty"`
	// Example: Foo Marketplace
	Affiliation *string `json:"affiliation,omitempty" gtag:"af,omitempty"`
	// Example: ChIJIQBpAG2ahYAR_6128GcTUEo
	LocationID *string `json:"location_id,omitempty" gtag:"lo,omitempty"`
	// The name of the promotional creative.
	// Example: summer_banner2
	CreativeName *string `json:"creative_name,omitempty" gtag:"cn,omitempty"`
	// The name of the promotional creative slot associated with the item.
	// Example: featured_app_1
	CreativeSlot *string `json:"creative_slot,omitempty" gtag:"cs,omitempty"`
	// The ID of the promotion associated with the item.
	// Example: P_12345
	PromotionID *string `json:"promotion_id,omitempty" gtag:"pi,omitempty"`
	// The name of the promotion associated with the item.
	// Example: Summer Sale
	PromotionName *string `json:"promotion_name,omitempty" gtag:"pn,omitempty"`
}
