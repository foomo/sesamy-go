package event

import (
	mp "github.com/foomo/sesamy-go/measurementprotocol"
	mpv2 "github.com/foomo/sesamy-go/measurementprotocol/v2"
)

type PageView struct{}

func (e *PageView) MarshalMPv2() (*mpv2.Event, error) {
	return &mpv2.Event{
		EventName: mp.Set(mpv2.EventNamePageView),
	}, nil
}
