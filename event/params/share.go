package params

// Share https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference/events#share
type Share struct {
	Method      string `json:"method,omitempty"`
	ContentType string `json:"content_type,omitempty"`
	ItemID      string `json:"item_id,omitempty"`
}
