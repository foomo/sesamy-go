package mpv2

type ConsentData struct {
	AdStorage              *Consent `json:"ad_storage,omitempty"`
	AdUserData             *Consent `json:"ad_user_data,omitempty"`
	AdPersonalization      *Consent `json:"ad_personalization,omitempty"`
	AnalyticsStorage       *Consent `json:"analytics_storage,omitempty"`
	FunctionalityStorage   *Consent `json:"functionality_storage,omitempty"`
	PersonalizationStorage *Consent `json:"personalization_storage,omitempty"`
	SecurityStorage        *Consent `json:"security_storage,omitempty"`
}
