package gtag

type Campaign struct {
	// Campaign Medium ( utm_medium ), this will override the current values read from the url
	// Example: cpc
	CampaignMedium *string `json:"campaign_medium,omitempty" gtag:"cm,omitempty"`
	// Campaign Source ( utm_source ), this will override the current values read from the url
	// Example: google
	CampaignSource *string `json:"campaign_source,omitempty" gtag:"cs,omitempty"`
	// Campaign Name ( utm_campaign ), this will override the current values read from the url
	// Example: cpc
	CampaignName *string `json:"campaign_name,omitempty" gtag:"cn,omitempty"`
	// Campaign Content ( utm_content ), this will override the current values read from the url
	// Example: big banner
	CampaignContent *string `json:"campaign_content,omitempty" gtag:"cc,omitempty"`
	// Campaign Term ( utm_term ), this will override the current values read from the url
	// Example: summer
	CampaignTerm *string `json:"campaign_term,omitempty" gtag:"ck,omitempty"`
	// Campaign Creative Format ( utm_creative_format ), this will override the current values read from the url
	// Example: native
	CampaignCreativeFormat *string `json:"campaign_creative_format,omitempty" gtag:"ccf,omitempty"`
	// Campaign Marketing Tactic ( utm_marketing_tactic ), this will override the current values read from the url
	// Example: prospecting
	CampaignMarketingTactic *string `json:"campaign_marketing_tactic,omitempty" gtag:"cmt,omitempty"`
	// Random Number used to Dedupe gclid
	// Example: 342342343
	// GclidDeduper *string `json:"gclid_deduper,omitempty" gtag:"_rnd,omitempty"`
}
