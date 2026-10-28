package params

// EmarsysView holds the parameters of the emarsys_view event: the viewed items
// and the Emarsys page view ID.
type EmarsysView[I any] struct {
	Items      []I    `json:"items,omitempty"`
	PageViewID string `json:"page_view_id,omitempty" dlv:"emarsys.page_view_id"`
}
