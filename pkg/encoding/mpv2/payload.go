package mpv2

import (
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type Payload[P any] struct {
	ClientID        string `json:"client_id,omitempty"`
	UserID          string `json:"user_id,omitempty"`
	TimestampMicros int64  `json:"timestamp_micros,omitempty"`
	// Reserved user property names: https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference?client_type=gtag#reserved_user_property_names
	UserProperties     map[string]any    `json:"user_properties,omitempty"`
	Consent            *ConsentData      `json:"consent,omitempty"`
	NonPersonalizedAds bool              `json:"non_personalized_ads,omitempty"`
	Events             []sesamy.Event[P] `json:"events,omitempty"`
	UserData           *UserData         `json:"user_data,omitempty"`
	DebugMode          bool              `json:"debug_mode,omitempty"`
}
