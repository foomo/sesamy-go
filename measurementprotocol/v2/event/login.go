package event

import (
	mp "github.com/foomo/sesamy/measurementprotocol"
	mpv2 "github.com/foomo/sesamy/measurementprotocol/v2"
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
	return &mpv2.Event{
		EventParameter: map[string]string{
			mpv2.EventParameterMethod.String(): *mp.SetString(e.Method),
		},
	}
}
