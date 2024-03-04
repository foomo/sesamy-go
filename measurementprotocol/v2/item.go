package v2
/**
  promotion_id: "pi",
  promotion_name: "pn",
  creative_name: "cn",
  creative_slot: "cs",
 */
type Item struct {
	// Exmaple: 12345
	ID *string `json:"id,omitempty" mapstructure:"id,omitempty"`
	// Example: Stan and Friends Tee
	Name *string `json:"nm,omitempty" mapstructure:"nm,omitempty"`
	// Exmaple: Google
	Brand *string `json:"br,omitempty" mapstructure:"br,omitempty"`
	// Exmaple: men
	CategoryHierarchy1 *string `json:"ca,omitempty" mapstructure:"ca,omitempty"`
	// Exmaple: t-shirts
	CategoryHierarchy2 *string `json:"c2,omitempty" mapstructure:"c2,omitempty"`
	// Exmaple: men
	CategoryHierarchy3 *string `json:"c3,omitempty" mapstructure:"c3,omitempty"`
	// Exmaple: men
	CategoryHierarchy4 *string `json:"c4,omitempty" mapstructure:"c4,omitempty"`
	// Exmaple: men
	CategoryHierarchy5 *string `json:"c5,omitempty" mapstructure:"c5,omitempty"`
	// Exmaple: Yellow
	Variant *string `json:"va,omitempty" mapstructure:"va,omitempty"`
	// Exmaple: 123.45
	Price *string `json:"pr,omitempty" mapstructure:"pr,omitempty"`
	// Exmaple: 1
	Quantity *string `json:"qt,omitempty" mapstructure:"qt,omitempty"`
	// Exmaple: 50%OFF
	Coupon *string `json:"cp,omitempty" mapstructure:"cp,omitempty"`
	// Exmaple: cross-selling: mens
	ListName *string `json:"ln,omitempty" mapstructure:"ln,omitempty"`
	// Exmaple: 10
	ListPosition *string `json:"lp,omitempty" mapstructure:"lp,omitempty"`
	// Exmaple: id-mens-123
	ListID *string `json:"li,omitempty" mapstructure:"li,omitempty"`
	// Exmaple: 10.00
	Discount *string `json:"ds,omitempty" mapstructure:"ds,omitempty"`
	// Exmaple: Foo Marketplace
	Affiliation *string `json:"af,omitempty" mapstructure:"af,omitempty"`
	// Example: ChIJIQBpAG2ahYAR_6128GcTUEo
	LocationID *string  `json:"lo,omitempty" mapstructure:"lo,omitempty"`
}
