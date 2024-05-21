package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type CampaignDetails sesamy2.Event[params.CampaignDetails]

func NewCampaignDetails(p params.CampaignDetails) CampaignDetails {
	return CampaignDetails(sesamy2.NewEvent(sesamy2.EventNameCampaignDetails, p))
}
