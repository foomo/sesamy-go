package mpv2

type ConsentData struct {
	AdUserData        *Consent `json:"ad_user_data,omitempty"`
	AdPersonalization *Consent `json:"ad_personalization,omitempty"`
}
