package params

// PostScore https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference/events#post_score
type PostScore struct {
	Score     int    `json:"score,omitempty"`
	Level     int    `json:"level,omitempty"`
	Character string `json:"character,omitempty"`
}
