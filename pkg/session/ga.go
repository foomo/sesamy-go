package session

import (
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

func ParseGAClientID(r *http.Request) (string, error) {
	cookie, err := r.Cookie("_ga")
	if err != nil {
		return "", errors.Wrap(err, "failed to retrieve _ga cookie")
	}

	parts := strings.Split(cookie.Value, ".")

	// validate
	if !strings.HasPrefix(cookie.Value, "GA1.1") || len(parts) < 4 {
		return "", errors.New("invalid _ga cookie value")
	}

	return parts[2] + "." + parts[3], nil
}

func ParseGASessionID(r *http.Request, id string) (string, error) {
	cookie, err := r.Cookie("_ga_" + id)
	if err != nil {
		return "", errors.Wrap(err, "failed to retrieve _ga cookie")
	}

	parts := strings.Split(cookie.Value, ".")

	// validate
	if !strings.HasPrefix(cookie.Value, "GS1.1") || len(parts) < 3 {
		return "", errors.New("invalid _ga cookie value")
	}

	return parts[2], nil
}
