package v2

type SST struct {
	// Example: 1
	ADR *string `json:"adr,omitempty" mapstructure:"adr,omitempty"`
	// Example: 1---
	USPrivacy *string `json:"us_privacy,omitempty" mapstructure:"us_privacy,omitempty"`
	// Example: 542231386.1709295522
	RND *string `json:"rnd,omitempty" mapstructure:"rnd,omitempty"`
	// Example: google.de
	ETLD *string `json:"etld,omitempty" mapstructure:"etld,omitempty"`
	// Example: region1
	GCSub *string `json:"gcsub,omitempty" mapstructure:"gcsub,omitempty"`
	// Example: DE
	UC *string `json:"uc,omitempty" mapstructure:"uc,omitempty"`
	// Example: 1708250245344
	TFT *string `json:"tft,omitempty" mapstructure:"tft,omitempty"`
	// Example: 13l3l3l3l1
	GCD *string `json:"gcd,omitempty" mapstructure:"gcd,omitempty"`
}
