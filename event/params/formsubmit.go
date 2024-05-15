package params

type FormSubmit struct {
	FormID          string `json:"form_id,omitempty"`
	FormName        string `json:"form_name,omitempty"`
	FormDestination string `json:"form_destination,omitempty"`
	FormSubmitText  string `json:"form_submit_text,omitempty"`
}
