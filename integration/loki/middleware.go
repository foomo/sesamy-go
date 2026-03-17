package loki

import (
	"net/http"

	"github.com/foomo/sesamy-go/pkg/encoding/gtag"
	"github.com/foomo/sesamy-go/pkg/encoding/gtagencode"
	"github.com/foomo/sesamy-go/pkg/encoding/mpv2"
	gtaghttp "github.com/foomo/sesamy-go/pkg/http/gtag"
	mpv2http "github.com/foomo/sesamy-go/pkg/http/mpv2"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func GTagMiddleware(loki *Loki) gtaghttp.Middleware {
	return func(next gtaghttp.MiddlewareHandler) gtaghttp.MiddlewareHandler {
		return func(l *zap.Logger, w http.ResponseWriter, r *http.Request, payload *gtag.Payload) error {
			err := next(l, w, r, payload)
			if err != nil {
				// encode to mpv2
				var mpv2Payload mpv2.Payload[any]
				if err := gtagencode.MPv2(*payload, &mpv2Payload); err != nil {
					return errors.Wrap(err, "failed to encode gtag to mpv2")
				}
				loki.Write(mpv2Payload)
			}
			return err
		}
	}
}

func MPv2Middleware(loki *Loki) mpv2http.Middleware {
	return func(next mpv2http.MiddlewareHandler) mpv2http.MiddlewareHandler {
		return func(l *zap.Logger, w http.ResponseWriter, r *http.Request, payload *mpv2.Payload[any]) error {
			err := next(l, w, r, payload)
			if err != nil {
				loki.Write(*payload)
			}
			return err
		}
	}
}
