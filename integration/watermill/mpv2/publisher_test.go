package mpv2_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/foomo/sesamy-go/integration/watermill/mpv2"
	encoding "github.com/foomo/sesamy-go/pkg/encoding/mpv2"
	"github.com/foomo/sesamy-go/pkg/event"
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

func TestPublisher(t *testing.T) {
	l := zaptest.NewLogger(t)

	var done atomic.Bool
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		out, err := io.ReadAll(r.Body)
		assert.NoError(t, err)

		expected := `{"client_id":"C123456","user_id":"U123456","timestamp_micros":1727701064057701,"events":[{"name":"page_view","params":{"debug_mode":true,"session_id":"S123456","engagement_time_msec":100,"page_title":"Home","page_location":"https://foomo.org"}}]}`
		if !assert.JSONEq(t, expected, string(out)) {
			fmt.Println(string(out))
		}
		done.Store(true)
	}))

	p := mpv2.NewPublisher(l, s.URL)

	payload := encoding.Payload[params.PageView]{
		ClientID:        "C123456",
		UserID:          "U123456",
		TimestampMicros: 1727701064057701,
		UserProperties:  nil,
		Consent:         nil,
		Events: []sesamy.Event[params.PageView]{
			event.NewPageView(params.PageView{
				MPv2: &params.MPv2{
					DebugMode:          true,
					SessionID:          "S123456",
					EngagementTimeMSec: 100,
				},
				PageTitle:    "Home",
				PageLocation: "https://foomo.org",
			}),
		},
		UserData: nil,
	}
	jsonPayload, err := json.Marshal(payload)
	require.NoError(t, err)

	msg := message.NewMessage(watermill.NewUUID(), jsonPayload)

	err = p.Publish("foo", msg)
	require.NoError(t, err)

	assert.Eventually(t, done.Load, time.Second, time.Millisecond)
}
