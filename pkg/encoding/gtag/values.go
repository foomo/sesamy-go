package gtag

import (
	"net/url"
	"strings"
)

// EncodeValues percent-encodes values as a query string, using %20 for spaces.
// NOTE: `richsstsse` seems to be last parameter in the query to let's ensure it stays that way
// The given values are not modified.
func EncodeValues(values url.Values) string {
	richsstsse := values.Has("richsstsse")
	if richsstsse {
		rest := make(url.Values, len(values))
		for k, v := range values {
			if k != "richsstsse" {
				rest[k] = v
			}
		}
		values = rest
	}

	ret := strings.ReplaceAll(values.Encode(), "+", "%20")

	if richsstsse {
		if ret != "" {
			ret += "&"
		}
		ret += "richsstsse"
	}
	return ret
}
