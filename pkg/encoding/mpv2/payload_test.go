package mpv2_test

import (
	"encoding/json"
	"testing"

	"github.com/foomo/sesamy-go/pkg/encoding/mpv2"
	"github.com/foomo/sesamy-go/pkg/event"
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPayload(t *testing.T) {
	t.Parallel()
	v := mpv2.Payload[params.PageView]{
		ClientID:        "C123456",
		UserID:          "U123456",
		TimestampMicros: 1727701064057701,
		UserProperties:  nil,
		Consent:         nil,
		Events: []sesamy.Event[params.PageView]{
			event.NewPageView(params.PageView{
				PageTitle:    "Home",
				PageLocation: "https://foomo.org",
			}),
		},
		UserData:           nil,
		DebugMode:          true,
		SessionID:          "S123456",
		EngagementTimeMSec: 100,
	}

	out, err := json.Marshal(v)
	require.NoError(t, err)
	expected := `{"debug_mode":true,"session_id":"S123456","engagement_time_msec":100,"client_id":"C123456","user_id":"U123456","timestamp_micros":1727701064057701,"events":[{"name":"page_view","params":{"page_title":"Home","page_location":"https://foomo.org"}}]}`
	assert.JSONEq(t, expected, string(out))

	var in mpv2.Payload[params.PageView]
	err = json.Unmarshal(out, &in)
	require.NoError(t, err)
	assert.Equal(t, v, in)
}
