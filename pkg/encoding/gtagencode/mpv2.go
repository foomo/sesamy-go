package gtagencode

import (
	"encoding/json"
	"fmt"
	"maps"
	"strconv"

	"github.com/foomo/sesamy-go/pkg/encoding/gtag"
	"github.com/foomo/sesamy-go/pkg/encoding/mpv2"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

func MPv2(source gtag.Payload, target any) error {
	var sourceData map[string]any

	out, err := json.Marshal(source)
	if err != nil {
		return errors.Wrap(err, "failed to marshal source")
	}
	if err = json.Unmarshal(out, &sourceData); err != nil {
		return errors.Wrap(err, "failed to unmarshal source")
	}

	// transform map to match mpv2 format
	targetData := map[string]any{
		"client_id":            source.ClientID,
		"user_id":              source.UserID,
		"session_id":           source.SessionID,
		"non_personalized_ads": source.NonPersonalizedAds,
		"debug_mode":           source.IsDebug,
	}

	// consent
	targetConsentData := map[string]any{
		"ad_storage":         mpv2.ConsentText(source.AdStorage()),
		"ad_user_data":       mpv2.ConsentText(source.AdUserData()),
		"ad_personalization": mpv2.ConsentText(source.AdPersonalization()),
		"analytics_storage":  mpv2.ConsentText(source.AnalyticsStorage()),
	}
	targetData["consent"] = targetConsentData

	// combine user properties
	targetUserProperties := map[string]any{}
	if node, ok := sourceData["user_property"].(map[string]string); ok {
		for s, s2 := range node {
			targetUserProperties[s] = s2
		}
	}
	if node, ok := sourceData["user_property_number"].(map[string]string); ok {
		for s, s2 := range node {
			targetUserProperties[s] = s2
		}
	}
	targetData["user_properties"] = targetUserProperties

	// transform event
	targetEventData := map[string]any{
		"name": source.EventName,
	}
	targetEventDataParams := map[string]any{}
	if value, ok := sourceData["document_title"]; ok {
		targetEventDataParams["page_title"] = value
	}
	if value, ok := sourceData["document_referrer"]; ok {
		targetEventDataParams["page_referrer"] = value
	}
	if value, ok := sourceData["document_location"]; ok {
		targetEventDataParams["page_location"] = value
	}
	if node, ok := sourceData["ecommerce"].(map[string]any); ok {
		maps.Copy(targetEventDataParams, node)
	}
	if node, ok := sourceData["event_parameter"].(map[string]any); ok {
		for s, s2 := range node {
			targetEventDataParams[s] = s2
		}
	}
	if node, ok := sourceData["event_parameter_number"].(map[string]any); ok {
		for s, s2 := range node {
			if value, err := strconv.ParseFloat(fmt.Sprintf("%s", s2), 64); err == nil {
				targetEventDataParams[s] = value
			} else {
				targetEventDataParams[s] = s2
			}
		}
	}
	targetEventData["params"] = targetEventDataParams
	targetData["events"] = []any{targetEventData}

	// encode map to target entity
	enc, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:               target,
		TagName:              "json",
		WeaklyTypedInput:     true,
		IgnoreUntaggedFields: true,
	})
	if err != nil {
		return errors.Wrap(err, "failed to create event encoder")
	}

	if err := enc.Decode(targetData); err != nil {
		return errors.Wrap(err, "failed to encode event")
	}

	return nil
}
