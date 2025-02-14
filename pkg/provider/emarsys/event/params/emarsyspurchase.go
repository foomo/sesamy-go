package params

type EmarsysPurchase[I any] struct {
	TransactionID string `json:"transaction_id,omitempty" dlv:"dataModel.transaction_id"`
	Items         []I    `json:"items,omitempty" dlv:"dataModel.items"`
	PageViewID    string `json:"page_view_id,omitempty" dlv:"emarsys.page_view_id"`
}
