package params

// LevelUp https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference/events#level_up
type LevelUp struct {
	Level     int    `json:"level,omitempty"`
	Character string `json:"character,omitempty"`
}
