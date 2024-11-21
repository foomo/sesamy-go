package mpv2

import (
	"encoding/json"
	"net/http"

	"github.com/foomo/sesamy-go/pkg/encoding/mpv2"
)

func Handler(w http.ResponseWriter, r *http.Request) *mpv2.Payload[any] {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return nil
	}

	// read request body
	var payload *mpv2.Payload[any]
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	// validate required fields
	if len(payload.Events) == 0 {
		http.Error(w, "missing events", http.StatusBadRequest)
		return nil
	}
	for _, event := range payload.Events {
		if event.Name == "" {
			http.Error(w, "missing event name", http.StatusBadRequest)
			return nil
		}
	}

	return payload
}
