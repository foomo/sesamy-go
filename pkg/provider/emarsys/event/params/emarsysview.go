package params

type EmarsysView[I any] struct {
	Item       I      `json:"item,omitempty"`
	PageViewID string `json:"page_view_id,omitempty" dlv:"emarsys.page_view_id"`
}
