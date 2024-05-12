package mpv2

type Consent struct {
	AdUserData        *string `json:"ad_user_data,omitempty"`
	AdPersonalization *string `json:"ad_personalization,omitempty"`
}
