package gtag

type SST struct {
	// Example: 1
	ADR *string `json:"adr,omitempty" gtag:"adr,omitempty"`
	// Example: 1---
	USPrivacy *string `json:"us_privacy,omitempty" gtag:"us_privacy,omitempty"`
	// Example: 542231386.1709295522
	RND *string `json:"rnd,omitempty" gtag:"rnd,omitempty"`
	// Example: google.de
	ETLD *string `json:"etld,omitempty" gtag:"etld,omitempty"`
	// Example: region1
	GCSub *string `json:"gcsub,omitempty" gtag:"gcsub,omitempty"`
	// Example: DE
	UC *string `json:"uc,omitempty" gtag:"uc,omitempty"`
	// Example: 1708250245344
	TFT *string `json:"tft,omitempty" gtag:"tft,omitempty"`
	// Example: 13l3l3l3l1
	GCD *string `json:"gcd,omitempty" gtag:"gcd,omitempty"`
	// Example: 0
	UDE *string `json:"ude,omitempty" gtag:"ude,omitempty"`
}
