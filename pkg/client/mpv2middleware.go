package client

import (
	"net/http"
	"strings"

	"github.com/foomo/sesamy-go/pkg/encoding/mpv2"
	"github.com/foomo/sesamy-go/pkg/session"
	"github.com/pkg/errors"
)

func MPv2MiddlewarSessionID(measurementID string) MPv2Middleware {
	measurementID = strings.Split(measurementID, "-")[1]
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

func MPv2MiddlewarDebugMode(next MPv2Handler) MPv2Handler {
	return func(r *http.Request, payload *mpv2.Payload[any]) error {
		if !payload.DebugMode {
			payload.DebugMode = session.IsGTMDebug(r)
		}
		return next(r, payload)
	}
}

func MPv2MiddlewareUserID(cookieName string) MPv2Middleware {
	return func(next MPv2Handler) MPv2Handler {
		return func(r *http.Request, payload *mpv2.Payload[any]) error {
			if payload.UserID == "" {
				value, err := r.Cookie(cookieName)
				if err != nil && !errors.Is(err, http.ErrNoCookie) {
					return err
				}
				payload.UserID = value.Value
			}
			return next(r, payload)
		}
	}
}
