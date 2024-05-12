package params

// LevelStart https://developers.google.com/analytics/devguides/collection/ga4/reference/events?client_type=gtag#level_start
type LevelStart struct {
	LevelName string `json:"level_name,omitempty"`
}
