package params

// EmarsysCart holds the parameters of the emarsys_cart event: the cart items
// and the Emarsys page view ID.
type EmarsysCart[I any] struct {
	Items      []I    `json:"items,omitempty"`
	PageViewID string `json:"page_view_id,omitempty" dlv:"emarsys.page_view_id"`
}
