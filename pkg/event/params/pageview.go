package params

// PageView https://developers.google.com/analytics/devguides/collection/ga4/views?client_type=gtag#manually_send_page_view_events
type PageView struct {
	*MPv2        `json:",inline"`
	PageTitle    string `json:"page_title,omitempty"`
	PageLocation string `json:"page_location,omitempty"`
}

type MPv2 struct {
	DebugMode          bool   `json:"debug_mode,omitempty"`
	SessionID          string `json:"session_id,omitempty"`
	EngagementTimeMSec int64  `json:"engagement_time_msec,omitempty"`
}
