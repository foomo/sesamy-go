package event

import (
	mp "github.com/foomo/sesamy-go/measurementprotocol"
	mpv2 "github.com/foomo/sesamy-go/measurementprotocol/v2"
)

/*
Login - Send this event to signify that a user has logged in to your website or app.

	gtag('event', 'login', {
		method: 'Google'
	});

Reference: https://developers.google.com/tag-platform/gtagjs/reference/events#login
*/
type Login struct {
	Method string
}

func (e *Login) MPv2() *mpv2.Event {
	eventParameter := map[string]string{}
	mp.AddStringMap(eventParameter, mpv2.EventParameterMethod.String(), mp.SetString(e.Method))
	return &mpv2.Event{
		EventName:      mp.Set(mpv2.EventNameLogin),
		EventParameter: mp.SetStringMap(eventParameter),
	}
}
