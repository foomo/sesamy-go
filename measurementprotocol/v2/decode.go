package v2

import (
	"net/url"
	"regexp"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

func Decode(input url.Values, output interface{}) error {
	data := Data{}

	// decode values
	for k, v := range input {
		// handle maps
		if ok, err := DecodeMapValue(k, v, data); err != nil {
			return err
		} else if ok {
			continue
		}

		// handle slices
		if ok, err := DecodeRegexValue(k, v, RegexProduct, data, ParameterItem); err != nil {
			return err
		} else if ok {
			continue
		}

		// default
		v, err := url.QueryUnescape(v[0])
		if err != nil {
			return err
		}
		data[k] = v
	}

	if err := mapstructure.WeakDecode(data, output); err != nil {
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
		if value, ok := data[parts[0]].(map[string]any); ok {
			v, err := url.QueryUnescape(v[0])
			if err != nil {
				return false, err
			}
			value[strings.Join(parts[1:], ".")] = v
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
		v, err := url.QueryUnescape(part[2:])
		if err != nil {
			return nil, err
		}
		ret[part[0:2]] = v
	}
	return ret, nil
}
