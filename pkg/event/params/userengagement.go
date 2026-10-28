package params

// UserEngagement holds the parameters of the automatically collected
// user_engagement event.
type UserEngagement struct {
	EngagementTimeMsec int64 `json:"engagement_time_msec,omitempty"`
}
