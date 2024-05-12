package params

// ViewSearchResults https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference/events#view_search_results
type ViewSearchResults struct {
	SearchTerm string `json:"search_term,omitempty"`
}
