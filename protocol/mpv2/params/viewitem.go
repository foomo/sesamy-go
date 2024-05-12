package params

// ViewItem https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference/events#view_item
type ViewItem[Item any] struct {
	Currency string  `json:"currency,omitempty"`
	Value    float64 `json:"value,omitempty"`
	Items    []Item  `json:"items,omitempty"`
}
