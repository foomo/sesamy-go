package params

// ViewItemList https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference/events#view_item_list
type ViewItemList[Item any] struct {
	ItemListID   string  `json:"item_list_id,omitempty"`
	ItemListName float64 `json:"item_list_name,omitempty"`
	Items        []Item  `json:"items,omitempty"`
}
