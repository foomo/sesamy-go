package params

// PageView https://developers.google.com/analytics/devguides/collection/ga4/views?client_type=gtag#manually_send_page_view_events
type PageView struct {
	PageTitle    string `json:"page_title,omitempty"`
	PageLocation string `json:"page_location,omitempty"`
}
