package params

// AdImpressionParams https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference/events#ad_impression
type AdImpressionParams struct {
	AdPlatform string `json:"ad_platform,omitempty" mapstructure:"ad_platform"`
	AdSource   string `json:"ad_source,omitempty" mapstructure:"ad_source"`
	AdFormat   string `json:"ad_format,omitempty" mapstructure:"ad_format"`
	AdUnitName string `json:"ad_unit_name,omitempty" mapstructure:"ad_unit_name"`
	Currency   string `json:"currency,omitempty" mapstructure:"currency"`
}
