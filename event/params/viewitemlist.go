package params

// ViewItemList https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference/events#view_item_list
type ViewItemList[I any] struct {
	ItemListID   string `json:"item_list_id,omitempty"`
	ItemListName string `json:"item_list_name,omitempty"`
	Items        []I    `json:"items,omitempty"`
}
