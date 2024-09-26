package client

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/foomo/sesamy-go/pkg/client"
	"github.com/foomo/sesamy-go/pkg/encoding/mpv2"
	"github.com/foomo/sesamy-go/pkg/provider/cookiebot"
)

func MPv2MiddlewarConsent(next client.MPv2Handler) client.MPv2Handler {
	return func(r *http.Request, payload *mpv2.Payload[any]) error {
		cookie, err := r.Cookie(cookiebot.CookieName)
		if err != nil || cookie.Value == "" {
			return next(r, payload)
		}

		data, err := base64.StdEncoding.DecodeString(cookie.Value)
		if err != nil {
			return next(r, payload)
		}

		var value cookiebot.Cookie
		if err := json.Unmarshal(data, &value); err != nil {
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
