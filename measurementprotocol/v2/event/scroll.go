package event

import (
	mp "github.com/foomo/sesamy-go/measurementprotocol"
	mpv2 "github.com/foomo/sesamy-go/measurementprotocol/v2"
)

type Scroll struct {
	PercentScrolled float64 `json:"percent_scrolled"`
}

func (e *Scroll) MarshalMPv2() (*mpv2.Event, error) {
	eventParameterNumber := map[string]string{}
	mp.AddStringMap(eventParameterNumber, mpv2.EventParameterNumberPercentScrolled.String(), mp.SetFloat64(e.PercentScrolled))
	return &mpv2.Event{
		EventName:            mp.Set(mpv2.EventNameScroll),
		EventParameterNumber: mp.SetStringMap(eventParameterNumber),
	}, nil
}
