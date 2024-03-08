package v2

import (
	"net/url"
)

// EncodeValues
// NOTE: `richsstsse` seems to be last parameter in the query to let's ensure it stays that way
func EncodeValues(values url.Values) string {
	var richsstsse bool
	if values.Has("richsstsse") {
		values.Del("richsstsse")
		richsstsse = true
	}

	ret := values.Encode()

	if richsstsse {
		ret += "&richsstsse"
	}
	return ret
}
