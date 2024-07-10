package params

// CampaignDetails https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference/events#campaign_details
type CampaignDetails struct {
	CampaignID string `json:"campaign_id,omitempty"`
	Campaign   string `json:"campaign,omitempty"`
	Source     string `json:"source,omitempty"`
	Medium     string `json:"medium,omitempty"`
	Term       string `json:"term,omitempty"`
	Content    string `json:"content,omitempty"`
}
