package v2

import (
	"net/http"
	"net/url"
	"strings"
)

func MiddlewareRichsstsse(next ClientHandler) ClientHandler {
	v := ""
	return func(r *http.Request, event *Event) error {
		event.Richsstsse = &v
		return next(r, event)
	}
}

func MiddlewareTrackingID(v string) ClientMiddleware {
	return func(next ClientHandler) ClientHandler {
		return func(r *http.Request, event *Event) error {
			event.TrackingID = &v
			return next(r, event)
		}
	}
}

func MiddlewarProtocolVersion(v string) ClientMiddleware {
	return func(next ClientHandler) ClientHandler {
		return func(r *http.Request, event *Event) error {
			event.ProtocolVersion = &v
			return next(r, event)
		}
	}
}

func MiddlewarIgnoreReferrer(v string) ClientMiddleware {
	return func(next ClientHandler) ClientHandler {
		return func(r *http.Request, event *Event) error {
			event.IgnoreReferrer = &v
			return next(r, event)
		}
	}
}

func MiddlewarDebug(next ClientHandler) ClientHandler {
	v := "1"
	return func(r *http.Request, event *Event) error {
		if value, _ := r.Cookie("gtm_debug"); value != nil {
			event.IsDebug = &v
		}
		return next(r, event)
	}
}

func MiddlewarClientID(next ClientHandler) ClientHandler {
	return func(r *http.Request, event *Event) error {
		if value, _ := r.Cookie("_ga"); value != nil {
			clientID := strings.TrimPrefix(value.Value, "GA1.1.")
			event.ClientID = &clientID
		}
		return next(r, event)
	}
}

func MiddlewarSessionID(trackingID string) ClientMiddleware {
	trackingID = strings.Split(trackingID, "-")[1]
	return func(next ClientHandler) ClientHandler {
		return func(r *http.Request, event *Event) error {
			if value, _ := r.Cookie("_ga_" + trackingID); value != nil {
				if value := strings.Split(value.Value, "."); len(value) > 3 {
					event.SessionID = &value[2]
				}
			}
			return next(r, event)
		}
	}
}

func MiddlewarDocument(next ClientHandler) ClientHandler {
	return func(r *http.Request, event *Event) error {
		if referrer, err := url.Parse(r.Referer()); err != nil {
			return err
		} else {
			location := referrer.RequestURI()
			event.DocumentLocation = &location
			event.DocumentHostname = &referrer.Host
		}
		return next(r, event)
	}
}
