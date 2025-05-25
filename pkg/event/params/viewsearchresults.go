package params

// ViewSearchResults https://developers.google.com/tag-platform/gtagjs/reference/events#view_search_results
type ViewSearchResults struct {
	SearchTerm string `json:"search_term,omitempty"`
}
