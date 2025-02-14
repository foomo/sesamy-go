package params

type EmarsysViewItemList struct {
	ItemListName string `json:"item_list_name,omitempty" dlv:"eventModel.item_list_name"`
	PageViewID   string `json:"page_view_id,omitempty" dlv:"emarsys.page_view_id"`
}
