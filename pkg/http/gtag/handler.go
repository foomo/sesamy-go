package gtag

import (
	"fmt"
	"io"
	"maps"
	"net/http"
	"net/url"
	"strings"

	"github.com/foomo/sesamy-go/pkg/encoding/gtag"
)

// Handler returns the first payload of the request, use HandlerBatch for batched requests.
func Handler(w http.ResponseWriter, r *http.Request) *gtag.Payload {
	payloads := HandlerBatch(w, r)
	if payloads == nil {
		return nil
	}

	return payloads[0]
}

// HandlerBatch returns one payload per event. gtag.js batches events as
// newline separated body lines, sharing the parameters from the query.
func HandlerBatch(w http.ResponseWriter, r *http.Request) []*gtag.Payload {
	query := r.URL.Query()
	lines := []string{""}

	switch r.Method {
	case http.MethodGet:
	case http.MethodPost:
		// read request body
		out, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to read body: %s", err.Error()), http.StatusInternalServerError)
			return nil
		}
		defer r.Body.Close()

		if body := strings.TrimSpace(string(out)); body != "" {
			lines = lines[:0]
			for line := range strings.SplitSeq(body, "\n") {
				if line = strings.TrimSpace(line); line != "" {
					lines = append(lines, line)
				}
			}
		}
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return nil
	}

	payloads := make([]*gtag.Payload, 0, len(lines))

	for _, line := range lines {
		// append body line to query
		values := maps.Clone(query)

		v, err := url.ParseQuery(line)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to parse extended url: %s", err.Error()), http.StatusInternalServerError)
			return nil
		}

		for s2, i := range v {
			values.Set(s2, i[0])
		}

		// unmarshal event
		var payload *gtag.Payload
		if err := gtag.Decode(values, &payload); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return nil
		}

		// validate
		if payload.EventName == nil || payload.EventName.String() == "" {
			http.Error(w, "missing event name", http.StatusBadRequest)
			return nil
		}

		payloads = append(payloads, payload)
	}

	return payloads
}
