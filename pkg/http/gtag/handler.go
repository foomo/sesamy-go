package gtag

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/foomo/sesamy-go/pkg/encoding/gtag"
)

func Handler(w http.ResponseWriter, r *http.Request) *gtag.Payload {
	var values url.Values

	switch r.Method {
	case http.MethodGet:
		values = r.URL.Query()
	case http.MethodPost:
		values = r.URL.Query()

		// read request body
		out, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to read body: %s", err.Error()), http.StatusInternalServerError)
			return nil
		}
		defer r.Body.Close()

		// append request body to query
		if len(out) > 0 {
			v, err := url.ParseQuery(string(out))
			if err != nil {
				http.Error(w, fmt.Sprintf("failed to parse extended url: %s", err.Error()), http.StatusInternalServerError)
				return nil
			}
			for s2, i := range v {
				values.Set(s2, i[0])
			}
		} else {
			values = r.URL.Query()
		}
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return nil
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

	return payload
}
