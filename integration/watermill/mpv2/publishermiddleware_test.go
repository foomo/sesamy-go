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

func TestPublisherMiddlewareIgnoreError(t *testing.T) {
	l := zaptest.NewLogger(t)

	var done atomic.Bool
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		done.Store(true)
	}))

	p := mpv2.NewPublisher(l, s.URL, mpv2.PublisherWithMiddlewares(mpv2.PublisherMiddlewareIgnoreError))

	payload := encoding.Payload[params.PageView]{
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

	jsonPayload, err := json.Marshal(payload)
	require.NoError(t, err)

	msg := message.NewMessage(watermill.NewUUID(), jsonPayload)

	require.NoError(t, p.Publish("foo", msg))

	assert.Eventually(t, done.Load, time.Second, 50*time.Millisecond)
}

func TestPublisherMiddlewareEventParams(t *testing.T) {
	l := zaptest.NewLogger(t)

	var done atomic.Bool
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		out, err := io.ReadAll(r.Body)
		assert.NoError(t, err)

		expected := `{"client_id":"C123456","user_id":"U123456","timestamp_micros":1727701064057701,"events":[{"name":"page_view","params":{"debug_mode":"1","engagement_time_msec":100,"page_location":"https://foomo.org","page_title":"Home","session_id":"S123456"}}]}`
		if !assert.JSONEq(t, expected, string(out)) {
			fmt.Println(string(out))
		}
		done.Store(true)
	}))

	p := mpv2.NewPublisher(l, s.URL, mpv2.PublisherWithMiddlewares(mpv2.PublisherMiddlewareEventParams))

	payload := encoding.Payload[params.PageView]{
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
	jsonPayload, err := json.Marshal(payload)
	require.NoError(t, err)

	msg := message.NewMessage(watermill.NewUUID(), jsonPayload)

	require.NoError(t, p.Publish("foo", msg))

	assert.Eventually(t, done.Load, time.Second, 50*time.Millisecond)
}
