package params

// SelectContent https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference/events#select_content
type SelectContent struct {
	ContentType string `json:"content_type,omitempty"`
	ContentID   string `json:"content_id,omitempty"`
}
