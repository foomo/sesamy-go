package params

type Scroll struct {
	PercentScrolled    int64 `json:"percent_scrolled,omitempty"`
	EngagementTimeMsec int64 `json:"engagement_time_msec,omitempty"`
}
