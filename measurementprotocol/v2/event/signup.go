package event

import (
	mp "github.com/foomo/sesamy-go/measurementprotocol"
	mpv2 "github.com/foomo/sesamy-go/measurementprotocol/v2"
)

/*
SignUp - Send this event to signify that a user has logged in to your website or app.

	gtag('event', 'sign_up', {
		method: 'Google'
	});

Reference: https://developers.google.com/tag-platform/gtagjs/reference/events#sign_up
*/
type SignUp struct {
	Method string `json:"method,omitempty"`
}

func (e *SignUp) MarshalMPv2() (*mpv2.Event, error) {
	eventParameter := map[string]string{}
	mp.AddStringMap(eventParameter, mpv2.EventParameterMethod.String(), mp.SetString(e.Method))
	return &mpv2.Event{
		EventName:      mp.Set(mpv2.EventNameSignUp),
		EventParameter: mp.SetStringMap(eventParameter),
	}, nil
}
