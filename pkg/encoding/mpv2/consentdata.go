package mpv2

// ConsentData holds the consent state per Google consent mode type.
// A nil field means the state is unknown.
type ConsentData struct {
	AdStorage              *Consent `json:"ad_storage,omitempty"`
	AdUserData             *Consent `json:"ad_user_data,omitempty"`
	AdPersonalization      *Consent `json:"ad_personalization,omitempty"`
	AnalyticsStorage       *Consent `json:"analytics_storage,omitempty"`
	FunctionalityStorage   *Consent `json:"functionality_storage,omitempty"`
	PersonalizationStorage *Consent `json:"personalization_storage,omitempty"`
	SecurityStorage        *Consent `json:"security_storage,omitempty"`
}
