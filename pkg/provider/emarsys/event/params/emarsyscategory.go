package params

type EmarsysCategory struct {
	ItemListName string `json:"item_list_name,omitempty"`
	PageViewID   string `json:"page_view_id,omitempty" dlv:"emarsys.page_view_id"`
}
