package params

// ScreenView https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference/events#screen_view
type ScreenView struct {
	ScreenClass string `json:"screen_class,omitempty"`
	ScreenName  string `json:"screen_name,omitempty"`
}
