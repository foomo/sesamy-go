package v2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"sort"
	"strings"

	"github.com/pkg/errors"
)

func Encode(input *Event) (url.Values, io.Reader, error) {
	var richsstsse bool
	// NOTE: `richsstsse` seems to be last parameter in the query to let's ensure it stays that way
	if input.Richsstsse != nil {
		richsstsse = true
		input.Richsstsse = nil
	}

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
	maxQueryLength := 2048 //
	if richsstsse {
		maxQueryLength -= len("&richsstsse")
	}
	for len(ret.Encode()) > maxQueryLength {
		for s, i := range ret {
			ret.Del(s)
			body = append(body, s+"="+i[0])
			break
		}
	}

	if richsstsse {
		ret.Add("richsstsse", "")
	}

	if len(body) > 0 {
		reader = bytes.NewReader([]byte(strings.Join(body, "&")))
	}

	return ret, reader, nil
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
	ret := make([]string, 0, len(keys))
	for _, k := range keys {
		ret = append(ret, k+fmt.Sprintf("%s", s[k]))
	}
	return strings.Join(ret, "~")
}
