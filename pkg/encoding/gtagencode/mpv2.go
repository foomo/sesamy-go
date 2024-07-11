package gtagencode

import (
	"encoding/json"
	"maps"

	"github.com/foomo/sesamy-go/pkg/encoding/gtag"
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
		"non_personalized_ads": source.NonPersonalizedAds,
		"debug_mode":           source.IsDebug,
	}
	if source.SST != nil && source.SST.TFT != nil {
		targetData["timestamp_micros"] = gtag.Get(source.SST.TFT) + "000"
	}

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
			targetEventDataParams[s] = s2
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
