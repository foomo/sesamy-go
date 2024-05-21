package params

import (
	"github.com/foomo/gostandards/iso4217"
)

// AdImpression https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference/events#ad_impression
type AdImpression struct {
	AdPlatform string           `json:"ad_platform,omitempty"`
	AdSource   string           `json:"ad_source,omitempty"`
	AdFormat   string           `json:"ad_format,omitempty"`
	AdUnitName string           `json:"ad_unit_name,omitempty"`
	Currency   iso4217.Currency `json:"currency,omitempty"`
	Value      float64          `json:"value,omitempty"`
}
