package gtag

import (
	"slices"
	"strings"
)

// See https://developers.google.com/tag-platform/security/concepts/consent-mode
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
	// Example: 1
	// DigitalMarketAct *string `json:"digital_market_act,omitempty" gtag:"dma,omitempty"`
	// Example: sypham
	// DigitalMarketActParameters *string `json:"digital_market_act_parameters,omitempty" gtag:"dma_cps,omitempty"`
	// Example: noapi | denied
}

// ------------------------------------------------------------------------------------------------
// ~ Public methods
// ------------------------------------------------------------------------------------------------

func (c Consent) AdStorage() bool {
	if c.GoogleConsentDefault != nil {
		gcd := strings.Split(*c.GoogleConsentDefault, "")
		if len(gcd) > 3 {
			return slices.Contains([]string{"l", "t", "r", "n", "u", "v"}, gcd[2])
		}
	}
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
	if c.GoogleConsentDefault != nil {
		gcd := strings.Split(*c.GoogleConsentDefault, "")
		if len(gcd) > 5 {
			return slices.Contains([]string{"l", "t", "r", "n", "u", "v"}, gcd[4])
		}
	}
	if c.GoogleConsentUpdate != nil {
		gcs := *c.GoogleConsentUpdate
		if strings.HasPrefix(gcs, "G1") && len(gcs) == 4 {
			return gcs[3:4] == "1"
		}
		return false
	}
	return true
}

func (c Consent) AdUserData() bool {
	if c.GoogleConsentDefault != nil {
		gcd := strings.Split(*c.GoogleConsentDefault, "")
		if len(gcd) > 7 {
			return slices.Contains([]string{"l", "t", "r", "n", "u", "v"}, gcd[6])
		}
	}
	return c.AdStorage()
}

func (c Consent) AdPersonalization() bool {
	if c.GoogleConsentDefault != nil {
		gcd := strings.Split(*c.GoogleConsentDefault, "")
		if len(gcd) > 9 {
			return slices.Contains([]string{"l", "t", "r", "n", "u", "v"}, gcd[8])
		}
	}
	return c.AdStorage()
}
