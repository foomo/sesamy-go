package params

// SelectItem https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference/events#select_item
type SelectItem[Item any] struct {
	ItemListID   string `json:"item_list_id,omitempty" tagging:"item_list_id"`
	ItemListName string `json:"item_list_name,omitempty" tagging:"item_list_name"`
	Items        []Item `json:"items,omitempty" tagging:"pr"`
}
