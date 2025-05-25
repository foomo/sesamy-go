package params

// PageView https://developers.google.com/tag-platform/gtagjs/reference/events#page_view
type PageView struct {
	ClientID     string `json:"client_id,omitempty"`
	Language     string `json:"language,omitempty"`
	PageEncoding string `json:"page_encoding,omitempty"`
	PageLocation string `json:"page_location,omitempty"`
	PageReferrer string `json:"page_referrer,omitempty"`
	PageTitle    string `json:"page_title,omitempty"`
	UserAgent    string `json:"user_agent,omitempty"`
}
