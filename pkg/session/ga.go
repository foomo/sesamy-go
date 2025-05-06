package session

import (
	"net/http"
	"regexp"

	"github.com/pkg/errors"
)

// see https://www.bbccss.com/explanation-of-cookie-values-used-by-ga4.html/comment-page-1#comment-6684
var (
	GA1Regex = regexp.MustCompile(`^GA1\.(\d+)\.(\d+)\.(\d+)$`)
	// GS1.<domain_level>.<session_id>.<session_count>.<engagement_session>.<timestamp>.<countdown>.<enhanced_client_id>
	GS1Regex = regexp.MustCompile(`^GS1\.(\d+)\.(\d+)\.(\d+)\.(\d+)\.(\d+)\.(\d+)\.(\d+)\.(\d+)$`)
	// GS2.<domain_level>.s<session_id>$o<session_count>$g<engagement_session>$t<timestamp>$j<countdown>$l<undefined>$h<enhanced_client_id>
	GS2Regex = regexp.MustCompile(`^GS2\.(\d+)\.s(\d+)\$o(\d+)\$g(\d+)\$t(\d+)\$j(\d+)\$l(\d+)\$h(\d+)$`)
)

func ParseGAClientID(r *http.Request) (string, error) {
	cookie, err := r.Cookie("_ga")
	if err != nil {
		return "", errors.Wrap(err, "failed to retrieve _ga cookie")
	}

	// validate
	if !GA1Regex.MatchString(cookie.Value) {
		return "", errors.New("invalid _ga cookie value")
	}

	parts := GA1Regex.FindStringSubmatch(cookie.Value)

	return parts[2] + "." + parts[3], nil
}

func ParseGASessionID(r *http.Request, id string) (string, error) {
	cookie, err := r.Cookie("_ga_" + id)
	if err != nil {
		return "", errors.Wrap(err, "failed to retrieve _ga cookie")
	}

	switch {
	case GS1Regex.MatchString(cookie.Value):
		return GS1Regex.FindStringSubmatch(cookie.Value)[2], nil
	case GS2Regex.MatchString(cookie.Value):
		return GS2Regex.FindStringSubmatch(cookie.Value)[2], nil
	default:
		return "", errors.Wrap(errors.New("invalid _ga cookie value"), cookie.Value)
	}
}

func ParseGASessionNumber(r *http.Request, id string) (string, error) {
	cookie, err := r.Cookie("_ga_" + id)
	if err != nil {
		return "", errors.Wrap(err, "failed to retrieve _ga cookie")
	}

	switch {
	case GS1Regex.MatchString(cookie.Value):
		return GS1Regex.FindStringSubmatch(cookie.Value)[3], nil
	case GS2Regex.MatchString(cookie.Value):
		return GS2Regex.FindStringSubmatch(cookie.Value)[3], nil
	default:
		return "", errors.Wrap(errors.New("invalid _ga cookie value"), cookie.Value)
	}
}
