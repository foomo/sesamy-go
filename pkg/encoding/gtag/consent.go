package gtag

import (
	"strings"
)

type Consent struct {
	// Current Google Consent Status. Format 'G1'+'AdsStorageBoolStatus'`+'AnalyticsStorageBoolStatus'
	// Example:  G101
	GoogleConsentStatus *string `json:"google_consent_status,omitempty" gtag:"gcs,omitempty"`
	// Will be added with the value "1" if the Google Consent has just been updated (wait_for_update setting on GTAG)
	// Example: 1
	GoogleConsentUpdate *string `json:"google_consent_update,omitempty" gtag:"gcu,omitempty"`
	// Documented values, 1 or 2, no more info on the meaning
	// Example: 2
	GoogleConsentUpdateType *string `json:"google_consent_update_type,omitempty" gtag:"gcut,omitempty"`
	// Will be added with the value "1" if the Google Consent had a default value before getting an update
	// Example: G111
	GoogleConsentDefault *string `json:"google_consent_default,omitempty" gtag:"gcd,omitempty"`
}

// ------------------------------------------------------------------------------------------------
// ~ Public methods
// ------------------------------------------------------------------------------------------------

func (c Consent) AdStorage() bool {
	if c.GoogleConsentUpdate != nil {
		gcs := *c.GoogleConsentUpdate
		if strings.HasPrefix(gcs, "G1") && len(gcs) == 4 {
			return gcs[2:3] == "1"
		}
		return false
	}
	return true
}

func (c Consent) AnalyticsStorage() bool {
	if c.GoogleConsentUpdate != nil {
		gcs := *c.GoogleConsentUpdate
		if strings.HasPrefix(gcs, "G1") && len(gcs) == 4 {
			return gcs[3:4] == "1"
		}
		return false
	}
	return true
}
