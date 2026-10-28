package params

// FormStart holds the parameters of the enhanced measurement form_start event.
// See https://support.google.com/analytics/answer/9216061.
type FormStart struct {
	FormID          string `json:"form_id,omitempty"`
	FormName        string `json:"form_name,omitempty"`
	FormDestination string `json:"form_destination,omitempty"`
}
