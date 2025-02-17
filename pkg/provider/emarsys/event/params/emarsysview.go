package params

type EmarsysView[I any] struct {
	Items      []I    `json:"items,omitempty" dlv:"dataModel.items"`
	PageViewID string `json:"page_view_id,omitempty" dlv:"emarsys.page_view_id"`
}
