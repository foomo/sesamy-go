package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/encoding/mpv2/params"
)

type CampaignDetails sesamy.Event[params.CampaignDetails]

func NewCampaignDetails(p params.CampaignDetails) CampaignDetails {
	return CampaignDetails(sesamy.NewEvent(sesamy.EventNameCampaignDetails, p))
}
