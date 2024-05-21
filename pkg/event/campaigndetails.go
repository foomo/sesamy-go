package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type CampaignDetails sesamy.Event[params.CampaignDetails]

func NewCampaignDetails(p params.CampaignDetails) CampaignDetails {
	return CampaignDetails(sesamy.NewEvent(sesamy.EventNameCampaignDetails, p))
}
