package client

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/foomo/sesamy-go/pkg/client"
	"github.com/foomo/sesamy-go/pkg/encoding/mpv2"
	"github.com/foomo/sesamy-go/pkg/provider/cookiebot"
	"go.uber.org/zap"
)

func MPv2MiddlewarConsent(l *zap.Logger) client.MPv2Middleware {
	return func(next client.MPv2Handler) client.MPv2Handler {
		return func(r *http.Request, payload *mpv2.Payload[any]) error {
			cookie, err := r.Cookie(cookiebot.CookieName)
			if errors.Is(err, http.ErrNoCookie) {
				return next(r, payload)
			} else if err != nil {
				l.With(zap.Error(err)).Warn("failed to retrieve cookie bot cookie")
				return next(r, payload)
			} else if cookie.Value == "" {
				l.With(zap.Error(err)).Warn("empty cookie bot cookie")
				return next(r, payload)
			}

			data, err := url.PathUnescape(cookie.Value)
			if err != nil {
				l.With(zap.Error(err)).Warn("failed to unescape cookie bot cookie")
				return next(r, payload)
			}

			var value cookiebot.Cookie
			if err := json.Unmarshal([]byte(data), &value); err != nil {
				l.With(zap.Error(err)).Warn("failed to unmarshal cookie bot cookie")
				return next(r, payload)
			}

			consent := func(b bool) *string {
				ret := "denied"
				if b {
					ret = "granted"
				}
				return &ret
			}

			payload.Consent = &mpv2.Consent{
				AdUserData:        consent(value.Marketing),
				AdPersonalization: consent(value.Statistics),
			}

			return next(r, payload)
		}
	}
}
