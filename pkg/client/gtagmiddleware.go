package client

import (
	"context"
	"net/http"

	"github.com/foomo/sesamy-go/pkg/encoding/gtag"
	"github.com/foomo/sesamy-go/pkg/session"
	"github.com/pkg/errors"
)

// GTagMiddlewareRichsstsse sets the richsstsse parameter, which requests a
// server-sent event response.
func GTagMiddlewareRichsstsse(next GTagHandler) GTagHandler {
	v := ""

	return func(r *http.Request, payload *gtag.Payload) error {
		payload.Richsstsse = &v
		return next(r, payload)
	}
}

// GTagMiddlewareTrackingID sets the tracking ID (tid) to v.
func GTagMiddlewareTrackingID(v string) GTagMiddleware {
	return func(next GTagHandler) GTagHandler {
		return func(r *http.Request, payload *gtag.Payload) error {
			payload.TrackingID = &v
			return next(r, payload)
		}
	}
}

// GTagMiddlewarProtocolVersion sets the protocol version (v) to v.
func GTagMiddlewarProtocolVersion(v string) GTagMiddleware {
	return func(next GTagHandler) GTagHandler {
		return func(r *http.Request, payload *gtag.Payload) error {
			payload.ProtocolVersion = &v
			return next(r, payload)
		}
	}
}

// GTagMiddlewarIsDebug sets the debug flag if the request has a gtm_debug cookie.
func GTagMiddlewarIsDebug(next GTagHandler) GTagHandler {
	v := "1"

	return func(r *http.Request, payload *gtag.Payload) error {
		if session.IsGTMDebug(r) {
			payload.IsDebug = &v
		}

		return next(r, payload)
	}
}

// GTagMiddlewarClientID sets the client ID from the _ga cookie.
// A missing cookie is ignored; an invalid cookie value returns an error.
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

// GTagMiddlewarWithoutCancel detaches the request context from cancellation of
// the incoming request, so the upstream request completes after the client
// disconnects.
func GTagMiddlewarWithoutCancel(next GTagHandler) GTagHandler {
	return func(r *http.Request, payload *gtag.Payload) error {
		return next(r.WithContext(context.WithoutCancel(r.Context())), payload)
	}
}

// GTagMiddlewarSessionID sets the session ID from the _ga_<suffix> cookie, where
// suffix is measurementID without its prefix (see [session.MeasurementIDSuffix]).
// A missing cookie is ignored; an invalid cookie value returns an error.
func GTagMiddlewarSessionID(measurementID string) GTagMiddleware {
	measurementID = session.MeasurementIDSuffix(measurementID)

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
