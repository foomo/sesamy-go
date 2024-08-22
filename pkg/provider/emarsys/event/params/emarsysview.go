package params

type EmarsysView[I any] struct {
	Items      []I    `json:"sitem,omitempty"`
	PageViewID string `json:"page_view_id,omitempty" dlv:"emarsys.page_view_id"`
}
