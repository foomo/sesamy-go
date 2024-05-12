package params

import (
	"encoding/json"
	"fmt"
	"maps"
	"strconv"

	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/protocol/gtag"
	"github.com/foomo/sesamy-go/protocol/mpv2"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

// SelectItem https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference/events#select_item
type SelectItem[Item any] struct {
	ItemListID   string `json:"item_list_id,omitempty" tagging:"item_list_id"`
	ItemListName string `json:"item_list_name,omitempty" tagging:"item_list_name"`
	Items        []Item `json:"items,omitempty" tagging:"pr"`
}

func NewSelectItem[P any](params P) sesamy.Event[P] {
	return sesamy.Event[P]{
		Name:   sesamy.EventNameSelectItem,
		Params: params,
	}
}

func Encode[P any](source mpv2.Payload[P], target *gtag.Payload) error {
	targetData := map[string]any{
		"client_id":            source.ClientID,
		"user_id":              source.UserID,
		"non_personalized_ads": source.NonPersonalizedAds,
	}

	{ // user_property
		targetUserProperty := map[string]any{}
		targetUserPropertyNumber := map[string]any{}
		for k, v := range source.UserProperties {
			if s, ok := v.(string); ok {
				if f, err := strconv.ParseFloat(s, 64); err == nil {
					targetUserPropertyNumber[k] = f
				} else {
					targetUserProperty[k] = v
				}
			} else {
				targetUserProperty[k] = fmt.Sprintf("%s", v)
			}
		}
		targetData["user_property"] = targetUserProperty
		targetData["user_property_number"] = targetUserPropertyNumber
	}

	sourceData := map[string]interface{}{}

	out, _ := json.Marshal(source.Events[0])
	_ = json.Unmarshal(out, &sourceData)

	{ // ecommerce
		targetData["event_name"] = sourceData["name"]

		if params, ok := sourceData["params"].(map[string]any); ok {
			targetEcommerceData := map[string]any{
				"currency":       params["currency"],
				"promotion_id":   params["promotion_id"],
				"promotion_name": params["promotion_name"],
				"location_id":    params["location_id"],
				"is_conversion":  params["is_conversion"],
			}
			delete(params, "currency")
			delete(params, "promotion_id")
			delete(params, "promotion_name")
			delete(params, "location_id")
			delete(params, "is_conversion")

			targetData["ecommerce"] = targetEcommerceData
			targetData["items"] = params["items"]
			delete(params, "items")

			{ // user_property
				targetEventProperty := map[string]any{}
				targetEventPropertyNumber := map[string]any{}
				for k, v := range params {
					if s, ok := v.(string); ok {
						if f, err := strconv.ParseFloat(s, 64); err == nil {
							targetEventPropertyNumber[k] = f
						} else {
							targetEventProperty[k] = v
						}
					} else {
						targetEventProperty[k] = fmt.Sprintf("%s", v)
					}
				}
				targetData["event_parameter"] = targetEventProperty
				targetData["event_parameter_number"] = targetEventPropertyNumber
			}
		}
	}

	// encode gtag event to map
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		WeaklyTypedInput:     true,
		Squash:               true,
		Result:               &target,
		TagName:              "json",
		IgnoreUntaggedFields: true,
	})
	if err != nil {
		return errors.Wrap(err, "failed to create event decoder")
	}

	if err := dec.Decode(targetData); err != nil {
		return errors.Wrap(err, "failed to decode event")
	}

	return nil
}

func Decode[P any](source gtag.Payload, target *mpv2.Payload[P]) error {
	var sourceData map[string]any
	// encode gtag event to map
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:               &sourceData,
		TagName:              "json",
		IgnoreUntaggedFields: true,
	})
	if err != nil {
		return errors.Wrap(err, "failed to create event decoder")
	}

	if err := dec.Decode(source); err != nil {
		return errors.Wrap(err, "failed to decode event")
	}

	// transorm map to match mpv2 format
	targetData := map[string]any{
		"client_id":            source.ClientID,
		"user_id":              source.UserID,
		"non_personalized_ads": source.NonPersonalizedAds,
		"timestamp_micros":     source.SST.TFT,
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
	if node, ok := sourceData["event_parameter"].(map[string]string); ok {
		for s, s2 := range node {
			targetEventDataParams[s] = s2
		}
	}
	if node, ok := sourceData["event_parameter_number"].(map[string]string); ok {
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
