package mpv2encode

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/foomo/sesamy-go/pkg/encoding/mpv2"
	"github.com/go-viper/mapstructure/v2"
	"github.com/pkg/errors"
)

// GTag converts the Measurement Protocol v2 payload source into gtag.js format
// and decodes the result into target, typically a pointer to a [gtag.Payload].
func GTag[P any](source mpv2.Payload[P], target any) error {
	if len(source.Events) == 0 {
		return errors.New("missing events")
	}

	targetData := map[string]any{
		"client_id": source.ClientID,
		"user_id":   source.UserID,
	}

	if source.SessionID != "" {
		targetData["session_id"] = source.SessionID
	}

	if source.DebugMode {
		targetData["is_debug"] = "1"
	}

	if c := source.Consent; c != nil {
		targetData["google_consent_status"] = "G1" + consentBit(c.AdStorage) + consentBit(c.AnalyticsStorage)
		targetData["google_consent_default"] = "13" + consentLetter(c.AdStorage) +
			"3" + consentLetter(c.AnalyticsStorage) +
			"3" + consentLetter(c.AdUserData) +
			"2" + consentLetter(c.AdPersonalization) + "5l1"
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
				targetUserProperty[k] = fmt.Sprintf("%v", v)
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
			targetData["document_title"] = params["page_title"]
			delete(params, "page_title")
			targetData["document_referrer"] = params["page_referrer"]
			delete(params, "page_referrer")
			targetData["document_location"] = params["page_location"]
			delete(params, "page_location")
			targetData["currency"] = params["currency"]
			delete(params, "currency")
			targetData["promotion_id"] = params["promotion_id"]
			delete(params, "promotion_id")
			targetData["promotion_name"] = params["promotion_name"]
			delete(params, "promotion_name")
			targetData["location_id"] = params["location_id"]
			delete(params, "location_id")
			targetData["is_conversion"] = params["is_conversion"]
			delete(params, "is_conversion")
			targetData["items"] = params["items"]
			delete(params, "items")
			{ // user_property
				targetEventProperty := map[string]any{}
				targetEventPropertyNumber := map[string]any{}

				for k, v := range params {
					switch t := v.(type) {
					case float64:
						targetEventPropertyNumber[k] = v
					case string:
						if f, err := strconv.ParseFloat(t, 64); err == nil {
							targetEventPropertyNumber[k] = f
						} else {
							targetEventProperty[k] = v
						}
					default:
						targetEventProperty[k] = fmt.Sprintf("%v", v)
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
		Result:               target,
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

// consentBit returns the gcs digit of c, treating an unset consent as granted.
func consentBit(c *mpv2.Consent) string {
	if c != nil && *c == mpv2.ConsentDenied {
		return "0"
	}

	return "1"
}

// consentLetter returns the gcd letter of c: l (not set), t (granted) or p (denied).
func consentLetter(c *mpv2.Consent) string {
	switch {
	case c == nil:
		return "l"
	case *c == mpv2.ConsentDenied:
		return "p"
	default:
		return "t"
	}
}
