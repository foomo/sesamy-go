package mpv2

import (
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference?client_type=gtag#payload_post_body
type Payload[P any] struct {
	ClientID           string            `json:"client_id,omitempty"`
	UserID             string            `json:"user_id,omitempty"`
	TimestampMicros    int64             `json:"timestamp_micros,omitempty"`
	UserProperties     map[string]any    `json:"user_properties,omitempty"`
	Consent            *ConsentData      `json:"consent,omitempty"`
	Events             []sesamy.Event[P] `json:"events,omitempty"`
	UserData           *UserData         `json:"user_data,omitempty"`
	DebugMode          bool              `json:"debug_mode,omitempty"`
	SessionID          string            `json:"session_id,omitempty"`
	EngagementTimeMSec int64             `json:"engagement_time_msec,omitempty"`
}
