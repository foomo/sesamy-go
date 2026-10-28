package session

import (
	"net/http"
)

// IsGTMDebug reports whether r carries the gtm_debug cookie that Google Tag
// Manager sets in debug mode.
func IsGTMDebug(r *http.Request) bool {
	_, err := r.Cookie("gtm_debug")
	return err == nil
}

// IsGTMPreview reports whether r carries the gtm_preview cookie that Google
// Tag Manager sets in preview mode.
func IsGTMPreview(r *http.Request) bool {
	_, err := r.Cookie("gtm_preview")
	return err == nil
}
