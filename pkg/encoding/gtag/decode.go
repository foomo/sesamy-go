package gtag

import (
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

func DecodeRequest(r http.Request, target any) error {
	u, err := url.ParseRequestURI(r.RequestURI)
	if err != nil {
		return errors.Wrap(err, "failed to parse request URI")
	}
	return Decode(u.Query(), target)
}

func DecodeQuery(input string, tarteg any) error {
	values, err := url.ParseQuery(input)
	if err != nil {
		return errors.Wrap(err, "failed to parse query")
	}
	return Decode(values, tarteg)
}

// Decode an incoming request into an Payload
func Decode(values url.Values, target any) error {
	data := Data{}

	// decode values
	for key, value := range values {
		// handle maps
		if ok, err := DecodeMapValue(key, value, data); err != nil {
			return errors.Wrap(err, "failed to decode map value")
		} else if ok {
			continue
		}

		// handle slices
		if ok, err := DecodeRegexValue(key, value, RegexProduct, data, ParameterItem); err != nil {
			return errors.Wrap(err, "failed to decode regex value")
		} else if ok {
			continue
		}

		// default
		data[key] = value[0]
	}

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		WeaklyTypedInput:     true,
		Result:               &target,
		TagName:              "gtag",
		IgnoreUntaggedFields: true,
		Squash:               true,
	})
	if err != nil {
		return errors.Wrap(err, "failed to weakly decode query")
	}

	if err := decoder.Decode(data); err != nil {
		return errors.Wrap(err, "failed to weakly decode query")
	}
	return nil
}

func DecodeMapValue(k string, v []string, data Data) (bool, error) {
	if strings.Contains(k, ".") {
		parts := strings.Split(k, ".")
		if _, ok := data[parts[0]]; !ok {
			data[parts[0]] = map[string]any{}
		}
		if value, ok := data[parts[0]].(map[string]any); ok && len(v) > 0 {
			val := v[0]
			// gracefully try to unescape value
			if out, err := url.QueryUnescape(val); err == nil {
				val = out
			}
			value[strings.Join(parts[1:], ".")] = val
		}
		return true, nil
	}
	return false, nil
}

// DecodeRegexValue e.g. `pr1=idSKU_123456` = map["pr"][]map["id"]="SKU_123456"
func DecodeRegexValue(k string, v []string, r *regexp.Regexp, data Data, key string) (bool, error) {
	if r.MatchString(k) {
		value, err := DecodeObjectValue(v[0])
		if err != nil {
			return false, err
		}
		if value != nil {
			v, ok := data[key].([]map[string]any)
			if !ok {
				v = []map[string]any{}
			}
			v = append(v, value)
			data[key] = v
			return true, nil
		}
	}
	return false, nil
}

// DecodeObjectValue e.g. `idSKU_123456` = map["id"]="SKU_123456"
func DecodeObjectValue(s string) (map[string]any, error) {
	if len(s) == 0 {
		return nil, nil //nolint:nilnil
	}
	ret := map[string]any{}
	for _, part := range strings.Split(s, "~") {
		val := part[2:]
		// gracefully try to unescape value
		if out, err := url.QueryUnescape(val); err == nil {
			val = out
		}
		ret[part[0:2]] = val
	}
	return ret, nil
}
