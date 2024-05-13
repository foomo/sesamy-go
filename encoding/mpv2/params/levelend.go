package params

// LevelEnd https://developers.google.com/analytics/devguides/collection/ga4/reference/events?client_type=gtag#level_end
type LevelEnd struct {
	LevelName string `json:"level_name,omitempty"`
	Success   bool   `json:"success,omitempty"`
}
