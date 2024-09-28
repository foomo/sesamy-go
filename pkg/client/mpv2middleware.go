package client

import (
	"net/http"

	"github.com/foomo/sesamy-go/pkg/encoding/mpv2"
	"github.com/foomo/sesamy-go/pkg/session"
	"github.com/pkg/errors"
)

func MPv2MiddlewarSessionID(trackingID string) MPv2Middleware {
	return func(next MPv2Handler) MPv2Handler {
		return func(r *http.Request, payload *mpv2.Payload[any]) error {
			value, err := session.ParseGASessionID(r, trackingID)
			if err != nil && !errors.Is(err, http.ErrNoCookie) {
				return errors.Wrap(err, "failed to parse client cookie")
			}
			if value != "" {
				payload.SessionID = value
			}
			return next(r, payload)
		}
	}
}

func MPv2MiddlewarClientID(next MPv2Handler) MPv2Handler {
	return func(r *http.Request, payload *mpv2.Payload[any]) error {
		value, err := session.ParseGAClientID(r)
		if err != nil && !errors.Is(err, http.ErrNoCookie) {
			return errors.Wrap(err, "failed to parse client cookie")
		}
		if value != "" {
			payload.ClientID = value
		}
		return next(r, payload)
	}
}

func MPv2MiddlewarDebugMode(next MPv2Handler) MPv2Handler {
	return func(r *http.Request, payload *mpv2.Payload[any]) error {
		if session.IsGTMDebug(r) {
			payload.DebugMode = true
		}
		return next(r, payload)
	}
}
