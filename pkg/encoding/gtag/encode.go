package gtag

import (
	"bytes"
	"fmt"
	"io"
	"net/url"
	"sort"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

func Encode(payload *Payload) (url.Values, io.Reader, error) {
	var richsstsse bool
	// NOTE: `richsstsse` seems to be last parameter in the query to let's ensure it stays that way
	if payload.Richsstsse != nil {
		richsstsse = true
		payload.Richsstsse = nil
	}

	remain := payload.Remain
	payload.Remain = nil

	data := Data{}
	var json = jsoniter.Config{
		TagKey:                        "gtag",
		EscapeHTML:                    false,
		MarshalFloatWith6Digits:       true, // will lose precession
		ObjectFieldMustBeSimpleString: true, // do not unescape object field
	}.Froze()

	jsonBytes, err := json.Marshal(&payload)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to marshall payload")
	}
	if err := json.Unmarshal(jsonBytes, &data); err != nil {
		return nil, nil, errors.Wrap(err, "failed to unmarshall payload")
	}

	for s, a := range remain {
		data[s] = a
	}

	ret := url.Values{}
	for k, v := range data {
		switch t := v.(type) {
		case []any:
			for key, value := range t {
				switch tt := value.(type) {
				case map[string]any:
					ret[fmt.Sprintf("%s%d", k, key+1)] = []string{EncodeObjectValue(tt)}
				default:
					panic("unhandled")
				}
			}
		case map[string]string:
			for key, value := range t {
				ret[fmt.Sprintf("%s.%s", k, key)] = []string{value}
			}
		case map[string]any:
			for key, value := range t {
				ret[fmt.Sprintf("%s.%s", k, key)] = []string{fmt.Sprintf("%v", value)}
			}
		case *string:
			ret[k] = []string{*t}
		case string:
			ret[k] = []string{t}
		case interface{ String() string }:
			ret[k] = []string{t.String()}
		case nil:
			continue
		default:
			panic("unhandled")
		}
	}

	var body []string
	var reader io.Reader
	maxQueryLength := 2048
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
