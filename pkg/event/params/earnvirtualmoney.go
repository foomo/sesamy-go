package params

// EarnVirtualMoney https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference/events#earn_virtual_currency
type EarnVirtualMoney struct {
	VirtualCurrencyName string  `json:"virtual_currency_name,omitempty"`
	Value               float64 `json:"value,omitempty"`
}
