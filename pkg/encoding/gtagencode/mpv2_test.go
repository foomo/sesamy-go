package gtagencode_test

import (
	"testing"

	"github.com/foomo/sesamy-go/pkg/encoding/gtag"
	"github.com/foomo/sesamy-go/pkg/encoding/gtagencode"
	"github.com/foomo/sesamy-go/pkg/encoding/mpv2"
	"github.com/foomo/sesamy-go/pkg/sesamy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMPv2(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		source  gtag.Payload
		want    mpv2.Payload[any]
		wantErr bool
	}{
		{
			name: "basic conversion",
			source: gtag.Payload{
				ClientID:           gtag.Set("test-client"),
				UserID:             gtag.Set("test-user"),
				SessionID:          gtag.Set("test-session"),
				NonPersonalizedAds: gtag.Set("1"),
				IsDebug:            gtag.Set("true"),
				EventName:          gtag.Set(sesamy.EventName("page_view")),
				DocumentTitle:      gtag.Set("Test Page"),
				DocumentLocation:   gtag.Set("https://test.com"),
				DocumentReferrer:   gtag.Set("https://referrer.com"),
			},
			want: mpv2.Payload[any]{
				ClientID:       "test-client",
				UserID:         "test-user",
				SessionID:      "test-session",
				DebugMode:      true,
				UserProperties: map[string]any{},
				Events: []sesamy.Event[any]{
					{
						Name: "page_view",
						Params: map[string]any{
							"page_title":    "Test Page",
							"page_location": "https://test.com",
							"page_referrer": "https://referrer.com",
						},
					},
				},
				Consent: &mpv2.ConsentData{
					AdStorage:         gtag.Set(mpv2.ConsentGranted),
					AdUserData:        gtag.Set(mpv2.ConsentGranted),
					AdPersonalization: gtag.Set(mpv2.ConsentGranted),
					AnalyticsStorage:  gtag.Set(mpv2.ConsentGranted),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var got mpv2.Payload[any]
			err := gtagencode.MPv2(tt.source, &got)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
