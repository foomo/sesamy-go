package event

import (
	mp "github.com/foomo/sesamy-go/measurementprotocol"
	mpv2 "github.com/foomo/sesamy-go/measurementprotocol/v2"
)

type Item struct {
	Affiliation   string  `json:"affiliation,omitempty"`
	Coupon        string  `json:"coupon,omitempty"`
	Discount      float64 `json:"discount,omitempty"`
	Index         int     `json:"index,omitempty"`
	ItemBrand     string  `json:"item_brand,omitempty"`
	ItemCategory  string  `json:"item_category,omitempty"`
	ItemCategory2 string  `json:"item_category2,omitempty"`
	ItemCategory3 string  `json:"item_category3,omitempty"`
	ItemCategory4 string  `json:"item_category4,omitempty"`
	ItemCategory5 string  `json:"item_category5,omitempty"`
	ItemID        string  `json:"item_id,omitempty"`
	ItemListName  string  `json:"item_list_name,omitempty"`
	ItemName      string  `json:"item_name,omitempty"`
	ItemVariant   string  `json:"item_variant,omitempty"`
	ItemListID    string  `json:"listId,omitempty"`
	LocationID    string  `json:"location_id,omitempty"`
	Price         string  `json:"price,omitempty"`
	Quantity      float64 `json:"quantity,omitempty"`
}

func (e *Item) MarshalMPv2() *mpv2.Item {
	return &mpv2.Item{
		ID:                 mp.SetString(e.ItemID),
		Name:               mp.SetString(e.ItemName),
		Brand:              mp.SetString(e.ItemBrand),
		CategoryHierarchy1: mp.SetString(e.ItemCategory),
		CategoryHierarchy2: mp.SetString(e.ItemCategory2),
		CategoryHierarchy3: mp.SetString(e.ItemCategory3),
		CategoryHierarchy4: mp.SetString(e.ItemCategory4),
		CategoryHierarchy5: mp.SetString(e.ItemCategory5),
		Price:              mp.SetString(e.Price),
		Quantity:           mp.SetFloat64(e.Quantity),
		Variant:            mp.SetString(e.ItemVariant),
		Coupon:             mp.SetString(e.Coupon),
		Discount:           mp.SetFloat64(e.Discount),
		ListName:           mp.SetString(e.ItemListName),
		ListID:             mp.SetString(e.ItemListID),
		ListPosition:       mp.SetInt(e.Index),
		Affiliation:        mp.SetString(e.Affiliation),
		LocationID:         mp.SetString(e.LocationID),
	}
}
