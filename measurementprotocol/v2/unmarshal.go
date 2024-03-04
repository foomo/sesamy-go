package v2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"regexp"
	"sort"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

var (
	RegexProduct = regexp.MustCompile(`pr([1-9]|[1-9][0-9]|1[0-9]{2}|200)`)
)

const (
	ParameterItem = "pr"
)

type Data map[string]any

func Marshal(input *Event) (url.Values, io.Reader, error) {

	a, err := json.Marshal(input)
	if err != nil {
		return nil, nil, err
	}

	data := Data{}
	if err := json.Unmarshal(a, &data); err != nil {
		return nil, nil, errors.Wrap(err, "failed to decode into map")
	}

	ret := url.Values{}
	for k, v := range data {
		switch t := v.(type) {
		case []interface{}:
			for key, value := range t {
				switch tt := value.(type) {
				case map[string]interface{}:
					ret[fmt.Sprintf("%s%d", k, key+1)] = []string{EncodeObjectValue(tt)}
				default:
					panic("unhandled")
				}
			}
		case map[string]string:
			for key, value := range t {
				ret[fmt.Sprintf("%s.%s", k, key)] = []string{value}
			}
		case map[string]interface{}:
			for key, value := range t {
				ret[fmt.Sprintf("%s.%s", k, key)] = []string{fmt.Sprintf("%v", value)}
			}
		case *string:
			ret[k] = []string{*t}
		case string:
			ret[k] = []string{t}
		default:
			panic("unhandled")
		}
	}

	var body []string
	var reader io.Reader
	for len(ret.Encode()) > 2048 {
		for s, i := range ret {
			ret.Del(s)
			body = append(body, s+"="+i[0])
			break
		}
	}
	if len(body) > 0 {
		reader = bytes.NewReader([]byte(strings.Join(body, "&")))
	}

	return ret, reader, nil
}


func UnmarshalURLValues(input url.Values, output interface{}) error {
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
			data[key] = value
			return true, nil
		}
	}
	return false, nil
}

// DecodeObjectValue e.g. `idSKU_123456` = map["id"]="SKU_123456"
func DecodeObjectValue(s string) (map[string]any, error) {
	if len(s) == 0 {
		return nil, nil
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

// EncodeObjectValue e.g. `idSKU_123456` = map["id"]="SKU_123456"
func EncodeObjectValue(s map[string]any) string {
	if len(s) == 0 {
		return ""
	}
	keys := make([]string, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var ret []string
	for _, k := range keys {
		ret = append(ret, k+fmt.Sprintf("%s", s[k]))
	}
	return strings.Join(ret, "~")
}
