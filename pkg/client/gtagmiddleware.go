package client

import (
	"net/http"
	"strings"

	"github.com/foomo/sesamy-go/pkg/encoding/gtag"
	"github.com/foomo/sesamy-go/pkg/session"
	"github.com/pkg/errors"
)

func GTagMiddlewareRichsstsse(next GTagHandler) GTagHandler {
	v := ""
	return func(r *http.Request, payload *gtag.Payload) error {
		payload.Richsstsse = &v
		return next(r, payload)
	}
}

func GTagMiddlewareTrackingID(v string) GTagMiddleware {
	return func(next GTagHandler) GTagHandler {
		return func(r *http.Request, payload *gtag.Payload) error {
			payload.TrackingID = &v
			return next(r, payload)
		}
	}
}

func GTagMiddlewarProtocolVersion(v string) GTagMiddleware {
	return func(next GTagHandler) GTagHandler {
		return func(r *http.Request, payload *gtag.Payload) error {
			payload.ProtocolVersion = &v
			return next(r, payload)
		}
	}
}

func GTagMiddlewarIsDebug(next GTagHandler) GTagHandler {
	v := "1"
	return func(r *http.Request, payload *gtag.Payload) error {
		if session.IsGTMDebug(r) {
			payload.IsDebug = &v
		}
		return next(r, payload)
	}
}

func GTagMiddlewarClientID(next GTagHandler) GTagHandler {
	return func(r *http.Request, payload *gtag.Payload) error {
		value, err := session.ParseGAClientID(r)
		if err != nil && !errors.Is(err, http.ErrNoCookie) {
			return errors.Wrap(err, "failed to parse client cookie")
		}
		if value != "" {
			payload.ClientID = &value
		}
		return next(r, payload)
	}
}

func GTagMiddlewarSessionID(measurementID string) GTagMiddleware {
	measurementID = strings.Split(measurementID, "-")[1]
	return func(next GTagHandler) GTagHandler {
		return func(r *http.Request, payload *gtag.Payload) error {
			value, err := session.ParseGASessionID(r, measurementID)
			if err != nil && !errors.Is(err, http.ErrNoCookie) {
				return errors.Wrap(err, "failed to parse session cookie")
			}
			if value != "" {
				payload.SessionID = &value
			}
			return next(r, payload)
		}
	}
}
