package params

type FormStart struct {
	FormID          string `json:"form_id,omitempty"`
	FormName        string `json:"form_name,omitempty"`
	FormDestination string `json:"form_destination,omitempty"`
}
