package params

type EmarsysCategory struct {
	Category   string `json:"category,omitempty"`
	PageViewID string `json:"page_view_id,omitempty" dlv:"emarsys.page_view_id"`
}
