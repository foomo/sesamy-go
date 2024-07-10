package params

// SelectItem https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference/events#select_item
type SelectItem[I any] struct {
	ItemListID   string `json:"item_list_id,omitempty" tagging:"item_list_id"`
	ItemListName string `json:"item_list_name,omitempty" tagging:"item_list_name"`
	Items        []I    `json:"items,omitempty" tagging:"pr"`
}
