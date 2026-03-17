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
				ClientID:           new("test-client"),
				UserID:             new("test-user"),
				SessionID:          new("test-session"),
				NonPersonalizedAds: new("1"),
				IsDebug:            new("true"),
				EventName:          new(sesamy.EventName("page_view")),
				DocumentTitle:      new("Test Page"),
				DocumentLocation:   new("https://test.com"),
				DocumentReferrer:   new("https://referrer.com"),
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
					AdStorage:         new(mpv2.ConsentGranted),
					AdUserData:        new(mpv2.ConsentGranted),
					AdPersonalization: new(mpv2.ConsentGranted),
					AnalyticsStorage:  new(mpv2.ConsentGranted),
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
