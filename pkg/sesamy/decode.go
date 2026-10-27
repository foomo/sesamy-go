package sesamy

import (
	"github.com/go-viper/mapstructure/v2"
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
