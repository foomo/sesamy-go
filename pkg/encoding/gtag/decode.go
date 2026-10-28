package gtag

import (
	"cmp"
	"maps"
	"net/http"
	"net/url"
	"regexp"
	"slices"
	"strings"

	"github.com/go-viper/mapstructure/v2"
	"github.com/pkg/errors"
)

// DecodeRequest decodes the query of r.RequestURI into target with [Decode].
func DecodeRequest(r http.Request, target any) error {
	u, err := url.ParseRequestURI(r.RequestURI)
	if err != nil {
		return errors.Wrap(err, "failed to parse request URI")
	}

	return Decode(u.Query(), target)
}

// DecodeQuery parses input as a URL query and decodes it into target with [Decode].
func DecodeQuery(input string, target any) error {
	values, err := url.ParseQuery(input)
	if err != nil {
		return errors.Wrap(err, "failed to parse query")
	}

	return Decode(values, target)
}

// Decode decodes gtag.js query values into target, typically a pointer to a
// [Payload]. Dotted keys become nested maps (see [DecodeMapValue]) and product
// keys such as pr1 become items (see [DecodeRegexValue]).
func Decode(values url.Values, target any) error {
	data := Data{}

	// decode values in length then lexical key order to keep items in pr1, pr2, …, pr10 order
	keys := slices.SortedFunc(maps.Keys(values), func(a, b string) int {
		return cmp.Or(cmp.Compare(len(a), len(b)), strings.Compare(a, b))
	})

	for _, key := range keys {
		value := values[key]
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

// DecodeMapValue stores a dotted key in data as a nested map, e.g. `ep.foo=bar`
// becomes data["ep"]["foo"]="bar". It reports whether k is a dotted key.
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

// DecodeRegexValue appends the object value of a key matching r to the slice
// data[key], e.g. `pr1=idSKU_123456` becomes data[key][]{"id": "SKU_123456"}.
// It reports whether a value was appended.
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

// DecodeObjectValue decodes a ~ separated list of two-letter keys and values,
// e.g. `idSKU_123456~nmShirt` becomes {"id": "SKU_123456", "nm": "Shirt"}.
// It returns nil for an empty s.
func DecodeObjectValue(s string) (map[string]any, error) {
	if len(s) == 0 {
		return nil, nil //nolint:nilnil
	}

	ret := map[string]any{}

	for part := range strings.SplitSeq(s, "~") {
		if part == "" {
			continue
		}

		if len(part) < 2 {
			return nil, errors.Errorf("invalid object value part %q", part)
		}

		val := part[2:]
		// gracefully try to unescape value
		if out, err := url.QueryUnescape(val); err == nil {
			val = out
		}

		ret[part[0:2]] = val
	}

	return ret, nil
}
