package client

import (
	"context"
	"net/http"

	"github.com/foomo/sesamy-go/pkg/encoding/mpv2"
	"github.com/foomo/sesamy-go/pkg/session"
	"github.com/pkg/errors"
)

// MPv2MiddlewarSessionID sets an empty session ID from the _ga_<suffix> cookie,
// where suffix is measurementID without its prefix (see
// [session.MeasurementIDSuffix]). A missing cookie is ignored; an invalid cookie
// value returns an error.
func MPv2MiddlewarSessionID(measurementID string) MPv2Middleware {
	measurementID = session.MeasurementIDSuffix(measurementID)

	return func(next MPv2Handler) MPv2Handler {
		return func(r *http.Request, payload *mpv2.Payload[any]) error {
			if payload.SessionID == "" {
				value, err := session.ParseGASessionID(r, measurementID)
				if err != nil && !errors.Is(err, http.ErrNoCookie) {
					return errors.Wrap(err, "failed to parse client cookie")
				}

				payload.SessionID = value
			}

			return next(r, payload)
		}
	}
}

// MPv2MiddlewarClientID sets an empty client ID from the _ga cookie.
// A missing cookie is ignored; an invalid cookie value returns an error.
func MPv2MiddlewarClientID(next MPv2Handler) MPv2Handler {
	return func(r *http.Request, payload *mpv2.Payload[any]) error {
		if payload.ClientID == "" {
			value, err := session.ParseGAClientID(r)
			if err != nil && !errors.Is(err, http.ErrNoCookie) {
				return errors.Wrap(err, "failed to parse client cookie")
			}

			payload.ClientID = value
		}

		return next(r, payload)
	}
}

// MPv2MiddlewarDebugMode enables debug mode if the request has a gtm_debug cookie.
func MPv2MiddlewarDebugMode(next MPv2Handler) MPv2Handler {
	return func(r *http.Request, payload *mpv2.Payload[any]) error {
		if !payload.DebugMode {
			payload.DebugMode = session.IsGTMDebug(r)
		}

		return next(r, payload)
	}
}

// MPv2MiddlewarWithoutCancel detaches the request context from cancellation of
// the incoming request, so the upstream request completes after the client
// disconnects.
func MPv2MiddlewarWithoutCancel(next MPv2Handler) MPv2Handler {
	return func(r *http.Request, payload *mpv2.Payload[any]) error {
		return next(r.WithContext(context.WithoutCancel(r.Context())), payload)
	}
}

// MPv2MiddlewareUserID sets an empty user ID from the cookie cookieName.
// A missing cookie is ignored.
func MPv2MiddlewareUserID(cookieName string) MPv2Middleware {
	return func(next MPv2Handler) MPv2Handler {
		return func(r *http.Request, payload *mpv2.Payload[any]) error {
			if payload.UserID == "" {
				value, err := r.Cookie(cookieName)
				if err != nil && !errors.Is(err, http.ErrNoCookie) {
					return err
				}

				if value != nil {
					payload.UserID = value.Value
				}
			}

			return next(r, payload)
		}
	}
}
