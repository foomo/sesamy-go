package sesamy

import (
	"github.com/go-viper/mapstructure/v2"
)

// Decode decodes input, typically a map, into output using json struct tags.
// It converts weakly typed values, e.g. a string to a number, and ignores
// untagged fields.
func Decode(input any, output any) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:               output,
		TagName:              "json",
		WeaklyTypedInput:     true,
		IgnoreUntaggedFields: true,
	})
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}
