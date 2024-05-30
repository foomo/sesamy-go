package session

import (
	"net/http"
)

func IsGTMDebug(r *http.Request) bool {
	_, err := r.Cookie("gtm_debug")
	return err == nil
}

func IsGTMPreview(r *http.Request) bool {
	_, err := r.Cookie("gtm_preview")
	return err == nil
}
