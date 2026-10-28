package params

// EmarsysPurchase holds the parameters of the emarsys_purchase event: the
// order and the Emarsys page view ID.
type EmarsysPurchase[I any] struct {
	TransactionID string `json:"transaction_id,omitempty"`
	Items         []I    `json:"items,omitempty"`
	PageViewID    string `json:"page_view_id,omitempty" dlv:"emarsys.page_view_id"`
}
