package gtag_test

import (
	"testing"

	"github.com/foomo/sesamy-go/pkg/encoding/gtag"
	"github.com/stretchr/testify/assert"
)

func TestConsent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name              string
		consent           gtag.Consent
		adStorage         bool
		analyticsStorage  bool
		adUserData        bool
		adPersonalization bool
	}{
		{
			name:              "unset",
			consent:           gtag.Consent{},
			adStorage:         true,
			analyticsStorage:  true,
			adUserData:        true,
			adPersonalization: true,
		},
		{
			name:              "gcs granted",
			consent:           gtag.Consent{GoogleConsentStatus: new("G111")},
			adStorage:         true,
			analyticsStorage:  true,
			adUserData:        true,
			adPersonalization: true,
		},
		{
			name:             "gcs analytics only",
			consent:          gtag.Consent{GoogleConsentStatus: new("G101")},
			analyticsStorage: true,
		},
		{
			name:      "gcs ads only",
			consent:   gtag.Consent{GoogleConsentStatus: new("G110")},
			adStorage: true, adUserData: true, adPersonalization: true,
		},
		{
			name:    "gcs denied",
			consent: gtag.Consent{GoogleConsentStatus: new("G100")},
		},
		{
			// gcu only flags an update and must not be read as a status
			name:              "gcu without gcs",
			consent:           gtag.Consent{GoogleConsentUpdate: new("1")},
			adStorage:         true,
			analyticsStorage:  true,
			adUserData:        true,
			adPersonalization: true,
		},
		{
			name:    "gcd takes precedence over gcs",
			consent: gtag.Consent{GoogleConsentStatus: new("G111"), GoogleConsentDefault: new("13p3p3p2p5l1")},
		},
		{
			name:              "gcd mixed",
			consent:           gtag.Consent{GoogleConsentDefault: new("13p3r3p2t5l1")},
			analyticsStorage:  true,
			adPersonalization: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.adStorage, tt.consent.AdStorage(), "ad_storage")
			assert.Equal(t, tt.analyticsStorage, tt.consent.AnalyticsStorage(), "analytics_storage")
			assert.Equal(t, tt.adUserData, tt.consent.AdUserData(), "ad_user_data")
			assert.Equal(t, tt.adPersonalization, tt.consent.AdPersonalization(), "ad_personalization")
		})
	}
}
