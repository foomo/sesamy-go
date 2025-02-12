package client

import (
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/foomo/sesamy-go/pkg/client"
	"github.com/foomo/sesamy-go/pkg/encoding/mpv2"
	"github.com/foomo/sesamy-go/pkg/provider/cookiebot"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
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

			data, err := url.QueryUnescape(cookie.Value)
			if err != nil {
				l.With(zap.Error(err), zap.String("value", cookie.Value)).Warn("failed to unescape cookie bot cookie")
				return next(r, payload)
			}

			var value cookiebot.Cookie
			if err := yaml.Unmarshal([]byte(strings.ReplaceAll(data, ":", ": ")), &value); err != nil {
				l.With(zap.Error(err), zap.String("value", data)).Warn("failed to unmarshal cookie bot cookie")
				return next(r, payload)
			}

			consent := func(b bool) *mpv2.Consent {
				ret := mpv2.ConsentDenied
				if b {
					ret = mpv2.ConsentGranted
				}
				return &ret
			}

			payload.Consent = &mpv2.ConsentData{
				AdStorage:              consent(value.Marketing),
				AdUserData:             consent(value.Marketing),
				AdPersonalization:      consent(value.Marketing),
				PersonalizationStorage: consent(value.Marketing),
				AnalyticsStorage:       consent(value.Statistics),
				FunctionalityStorage:   consent(value.Necessary),
				SecurityStorage:        consent(value.Necessary),
			}

			return next(r, payload)
		}
	}
}
