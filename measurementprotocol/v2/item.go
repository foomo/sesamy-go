package v2

/*
*

	promotion_id: "pi",
	promotion_name: "pn",
	creative_name: "cn",
	creative_slot: "cs",
*/
type Item struct {
	// Example: 12345
	ID *string `json:"id,omitempty" mapstructure:"id,omitempty"`
	// Example: Stan and Friends Tee
	Name *string `json:"nm,omitempty" mapstructure:"nm,omitempty"`
	// Example: Google
	Brand *string `json:"br,omitempty" mapstructure:"br,omitempty"`
	// Example: men
	CategoryHierarchy1 *string `json:"ca,omitempty" mapstructure:"ca,omitempty"`
	// Example: t-shirts
	CategoryHierarchy2 *string `json:"c2,omitempty" mapstructure:"c2,omitempty"`
	// Example: men
	CategoryHierarchy3 *string `json:"c3,omitempty" mapstructure:"c3,omitempty"`
	// Example: men
	CategoryHierarchy4 *string `json:"c4,omitempty" mapstructure:"c4,omitempty"`
	// Example: men
	CategoryHierarchy5 *string `json:"c5,omitempty" mapstructure:"c5,omitempty"`
	// Example: Yellow
	Variant *string `json:"va,omitempty" mapstructure:"va,omitempty"`
	// Example: 123.45
	Price *string `json:"pr,omitempty" mapstructure:"pr,omitempty"`
	// Example: 1
	Quantity *string `json:"qt,omitempty" mapstructure:"qt,omitempty"`
	// Example: 50%OFF
	Coupon *string `json:"cp,omitempty" mapstructure:"cp,omitempty"`
	// Example: cross-selling: mens
	ListName *string `json:"ln,omitempty" mapstructure:"ln,omitempty"`
	// Example: 10
	ListPosition *string `json:"lp,omitempty" mapstructure:"lp,omitempty"`
	// Example: id-mens-123
	ListID *string `json:"li,omitempty" mapstructure:"li,omitempty"`
	// Example: 10.00
	Discount *string `json:"ds,omitempty" mapstructure:"ds,omitempty"`
	// Example: Foo Marketplace
	Affiliation *string `json:"af,omitempty" mapstructure:"af,omitempty"`
	// Example: ChIJIQBpAG2ahYAR_6128GcTUEo
	LocationID *string `json:"lo,omitempty" mapstructure:"lo,omitempty"`
	// The name of the promotional creative.
	// Example: summer_banner2
	CreativeName *string `json:"cn,omitempty" mapstructure:"cn,omitempty"`
	// The name of the promotional creative slot associated with the item.
	// Example: featured_app_1
	CreativeSlot *string `json:"cs,omitempty" mapstructure:"cs,omitempty"`
	// The ID of the promotion associated with the item.
	// Example: P_12345
	PromotionID *string `json:"pi,omitempty" mapstructure:"pi,omitempty"`
	// The name of the promotion associated with the item.
	// Example: Summer Sale
	PromotionName *string `json:"pn,omitempty" mapstructure:"pn,omitempty"`
}
