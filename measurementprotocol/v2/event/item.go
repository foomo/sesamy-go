package event

import (
	mp "github.com/foomo/sesamy-go/measurementprotocol"
	mpv2 "github.com/foomo/sesamy-go/measurementprotocol/v2"
)

type Item struct {
	ID                 string
	Name               string
	Brand              string
	CategoryHierarchy1 string
	CategoryHierarchy2 string
	CategoryHierarchy3 string
	CategoryHierarchy4 string
	CategoryHierarchy5 string
	Price              string
	Quantity           float64
	Variant            string
	Coupon             string
	Discount           float64
	Index              int
	ListName           string
	ListID             string
	Affiliation        string
	LocationID         string
}

func (e *Item) MPv2() *mpv2.Item {
	return &mpv2.Item{
		ID:                 mp.SetString(e.ID),
		Name:               mp.SetString(e.Name),
		Brand:              mp.SetString(e.Brand),
		CategoryHierarchy1: mp.SetString(e.CategoryHierarchy1),
		CategoryHierarchy2: mp.SetString(e.CategoryHierarchy2),
		CategoryHierarchy3: mp.SetString(e.CategoryHierarchy3),
		CategoryHierarchy4: mp.SetString(e.CategoryHierarchy4),
		CategoryHierarchy5: mp.SetString(e.CategoryHierarchy5),
		Price:              mp.SetString(e.Price),
		Quantity:           mp.SetFloat64(e.Quantity),
		Variant:            mp.SetString(e.Variant),
		Coupon:             mp.SetString(e.Coupon),
		Discount:           mp.SetFloat64(e.Discount),
		ListName:           mp.SetString(e.ListName),
		ListID:             mp.SetString(e.ListID),
		ListPosition:       mp.SetInt(e.Index),
		Affiliation:        mp.SetString(e.Affiliation),
		LocationID:         mp.SetString(e.LocationID),
	}
}
