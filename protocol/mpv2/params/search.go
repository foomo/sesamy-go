package params

// Search https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference/events#search
type Search struct {
	SearchTerm string `json:"search_term,omitempty"`
}
