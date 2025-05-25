package params

// Exception https://developers.google.com/tag-platform/gtagjs/reference/events#parameters_12
type Exception struct {
	Description string `json:"description,omitempty"`
	Fatal       bool   `json:"fatal,omitempty"`
}
