package loki

import (
	"encoding/json"

	"github.com/foomo/sesamy-go/pkg/encoding/mpv2"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type Line struct {
	Name               sesamy.EventName `json:"name"`
	Params             any              `json:"params"`
	ClientID           string           `json:"client_id"`
	UserID             string           `json:"user_id,omitempty"`
	UserProperties     map[string]any   `json:"user_properties,omitempty"`
	Consent            *mpv2.Consent    `json:"consent,omitempty"`
	NonPersonalizedAds bool             `json:"non_personalized_ads,omitempty"`
	UserData           *mpv2.UserData   `json:"user_data,omitempty"`
	DebugMode          bool             `json:"debug_mode,omitempty"`
}

func (l *Line) Marshal() ([]byte, error) {
	return json.Marshal(l)
}
