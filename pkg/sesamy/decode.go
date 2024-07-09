package sesamy

import (
	"github.com/mitchellh/mapstructure"
)

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
