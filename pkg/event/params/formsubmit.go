package params

// FormSubmit holds the parameters of the enhanced measurement form_submit
// event. See https://support.google.com/analytics/answer/9216061.
type FormSubmit struct {
	FormID          string `json:"form_id,omitempty"`
	FormName        string `json:"form_name,omitempty"`
	FormDestination string `json:"form_destination,omitempty"`
	FormSubmitText  string `json:"form_submit_text,omitempty"`
}
