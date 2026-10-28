package params

// Scroll holds the parameters of the enhanced measurement scroll event.
// See https://support.google.com/analytics/answer/9216061.
type Scroll struct {
	PercentScrolled    int64 `json:"percent_scrolled,omitempty"`
	EngagementTimeMsec int64 `json:"engagement_time_msec,omitempty"`
}
