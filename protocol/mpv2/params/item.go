package params

type Item struct {
	Affiliation   string  `json:"affiliation,omitempty" tagging:"af"`
	Coupon        string  `json:"coupon,omitempty" tagging:"cp"`
	Discount      float64 `json:"discount,omitempty" tagging:"ds"`
	Index         int     `json:"index,omitempty" tagging:"lp"`
	ItemBrand     string  `json:"item_brand,omitempty" tagging:"br"`
	ItemCategory  string  `json:"item_category,omitempty" tagging:"ca"`
	ItemCategory2 string  `json:"item_category2,omitempty" tagging:"c2"`
	ItemCategory3 string  `json:"item_category3,omitempty" tagging:"c3"`
	ItemCategory4 string  `json:"item_category4,omitempty" tagging:"c4"`
	ItemCategory5 string  `json:"item_category5,omitempty" tagging:"c5"`
	ItemID        string  `json:"item_id,omitempty" tagging:"id"`
	ItemListID    string  `json:"listId,omitempty" tagging:"li"`
	ItemListName  string  `json:"item_list_name,omitempty" tagging:"ln"`
	ItemName      string  `json:"item_name,omitempty" tagging:"nm"`
	ItemVariant   string  `json:"item_variant,omitempty" tagging:"va"`
	LocationID    string  `json:"location_id,omitempty" tagging:"lo"`
	Price         string  `json:"price,omitempty" tagging:"pr"`
	Quantity      float64 `json:"quantity,omitempty" tagging:"qt"`
}
