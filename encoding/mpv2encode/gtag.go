package mpv2encode

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/foomo/sesamy-go/encoding/gtag"
	"github.com/foomo/sesamy-go/encoding/mpv2"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

func GTag[P any](source mpv2.Payload[P], target *gtag.Payload) error {
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

	sourceData := map[string]any{}

	out, err := json.Marshal(source.Events[0])
	if err != nil {
		return errors.Wrap(err, "failed to marshal event")
	}
	if err = json.Unmarshal(out, &sourceData); err != nil {
		return errors.Wrap(err, "failed to unmarshal source events")
	}

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
