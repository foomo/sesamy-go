package params

type Exception struct {
	Description string `json:"description,omitempty"`
	Fatal       bool   `json:"fatal,omitempty"`
}
