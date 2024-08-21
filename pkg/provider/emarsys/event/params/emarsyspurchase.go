package params

type EmarsysPurchase[I any] struct {
	Items      []I    `json:"items,omitempty"`
	PageViewID string `json:"page_view_id,omitempty" dlv:"emarsys.page_view_id"`
}
