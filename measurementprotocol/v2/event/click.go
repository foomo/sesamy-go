package event

import (
	mp "github.com/foomo/sesamy-go/measurementprotocol"
	mpv2 "github.com/foomo/sesamy-go/measurementprotocol/v2"
)

type Click struct {
	LinkID      string `json:"linkId,omitempty"`
	LinkURL     string `json:"link_url,omitempty"`
	LinkDomain  string `json:"link_domain,omitempty"`
	LinkClasses string `json:"link_classes,omitempty"`
	Outbound    bool   `json:"outbound,omitempty"`
}

func (e *Click) MarshalMPv2() (*mpv2.Event, error) {
	eventParameter := map[string]string{}
	mp.AddStringMap(eventParameter, mpv2.EventParameterLinkID.String(), mp.SetString(e.LinkID))
	mp.AddStringMap(eventParameter, mpv2.EventParameterLinkUrl.String(), mp.SetString(e.LinkURL))
	mp.AddStringMap(eventParameter, mpv2.EventParameterLinkDomain.String(), mp.SetString(e.LinkDomain))
	mp.AddStringMap(eventParameter, mpv2.EventParameterLinkClasses.String(), mp.SetString(e.LinkClasses))
	mp.AddStringMap(eventParameter, mpv2.EventParameterOutbound.String(), mp.SetBool(e.Outbound))
	return &mpv2.Event{
		EventName:      mp.Set(mpv2.EventNameClick),
		EventParameter: mp.SetStringMap(eventParameter),
	}, nil
}
