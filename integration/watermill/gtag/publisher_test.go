package gtag_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/foomo/sesamy-go/integration/watermill/gtag"
	encoding "github.com/foomo/sesamy-go/pkg/encoding/gtag"
	"github.com/foomo/sesamy-go/pkg/sesamy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

func TestPublisher(t *testing.T) {
	l := zaptest.NewLogger(t)

	var done atomic.Bool
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expected := `_dbg=1&cid=C123456&dl=https%3A%2F%2Ffoomo.org&dt=Home&en=add_to_cart&v=2&richsstsse`
		assert.Equal(t, expected, r.URL.RawQuery)
		done.Store(true)
	}))

	p := gtag.NewPublisher(l, s.URL)

	payload := encoding.Payload{
		Consent:              encoding.Consent{},
		Campaign:             encoding.Campaign{},
		ECommerce:            encoding.ECommerce{},
		ClientHints:          encoding.ClientHints{},
		ProtocolVersion:      encoding.Set("2"),
		TrackingID:           nil,
		GTMHashInfo:          nil,
		ClientID:             encoding.Set("C123456"),
		Richsstsse:           encoding.Set("1"),
		DocumentLocation:     encoding.Set("https://foomo.org"),
		DocumentTitle:        encoding.Set("Home"),
		DocumentReferrer:     nil,
		IsDebug:              encoding.Set("1"),
		EventName:            encoding.Set(sesamy.EventNameAddToCart),
		EventParameter:       nil,
		EventParameterNumber: nil,
		UserID:               nil,
		SessionID:            nil,
		UserProperty:         nil,
		UserPropertyNumber:   nil,
		NonPersonalizedAds:   nil,
		SST:                  nil,
		Remain:               nil,
	}
	jsonPayload, err := json.Marshal(payload)
	require.NoError(t, err)

	fmt.Println(string(jsonPayload))
	msg := message.NewMessage(watermill.NewUUID(), jsonPayload)

	require.NoError(t, p.Publish("foo", msg))

	assert.Eventually(t, done.Load, time.Second, 50*time.Millisecond)
}
