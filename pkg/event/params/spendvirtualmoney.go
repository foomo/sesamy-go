package params

// SpendVirtualCurrency https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference/events#spend_virtual_currency
type SpendVirtualCurrency struct {
	Value               float64 `json:"value,omitempty"`
	VirtualCurrencyName string  `json:"virtual_currency_name,omitempty"`
	ItemName            string  `json:"item_name,omitempty"`
}
