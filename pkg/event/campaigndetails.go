package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// CampaignDetails is the GA4 campaign_details event.
type CampaignDetails sesamy.Event[params.CampaignDetails]

// NewCampaignDetails returns a campaign_details event with the parameters p.
func NewCampaignDetails(p params.CampaignDetails) sesamy.Event[params.CampaignDetails] {
	return sesamy.NewEvent(sesamy.EventNameCampaignDetails, p)
}
