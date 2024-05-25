package client

import (
	"net/http"

	"github.com/foomo/sesamy-go/pkg/encoding/mpv2"
	"github.com/foomo/sesamy-go/pkg/session"
	"github.com/pkg/errors"
)

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
