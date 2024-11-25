package gtag_test

import (
	"context"
	"encoding/json"
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/foomo/sesamy-go/integration/watermill/gtag"
	encoding "github.com/foomo/sesamy-go/pkg/encoding/gtag"
	"github.com/foomo/sesamy-go/pkg/sesamy"
	"github.com/pperaltaisern/watermillzap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

func TestMessageHandler(t *testing.T) {
	l := zaptest.NewLogger(t)

	router, err := message.NewRouter(message.RouterConfig{}, watermillzap.NewLogger(l))
	require.NoError(t, err)
	defer router.Close()

	// Create pubSub
	pubSub := gochannel.NewGoChannel(
		gochannel.Config{},
		watermillzap.NewLogger(l),
	)

	var done atomic.Bool
	router.AddHandler("gtag", "in", pubSub, "out", pubSub, gtag.MessageHandler(func(payload *encoding.Payload, msg *message.Message) error {
		expected := `{"consent":{},"campaign":{},"ecommerce":{},"client_hints":{},"protocol_version":"2","client_id":"C123456","richsstsse":"1","document_location":"https://foomo.org","document_title":"Home","is_debug":"1","event_name":"add_to_cart"}`
		if !assert.JSONEq(t, expected, string(msg.Payload)) {
			fmt.Println(string(msg.Payload))
		}
		done.Store(true)
		return nil
	}))

	go func() {
		assert.NoError(t, router.Run(context.TODO()))
	}()
	assert.Eventually(t, router.IsRunning, time.Second, 50*time.Millisecond)

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

	msg := message.NewMessage(watermill.NewUUID(), jsonPayload)

	require.NoError(t, pubSub.Publish("in", msg))

	assert.Eventually(t, done.Load, time.Second, 50*time.Millisecond)
}

func TestMPv2MessageHandler(t *testing.T) {
	l := zaptest.NewLogger(t)

	router, err := message.NewRouter(message.RouterConfig{}, watermillzap.NewLogger(l))
	require.NoError(t, err)
	defer router.Close()

	// Create pubSub
	pubSub := gochannel.NewGoChannel(
		gochannel.Config{},
		watermillzap.NewLogger(l),
	)

	var done atomic.Bool
	router.AddHandler("gtag", "in", pubSub, "out", pubSub, gtag.MPv2MessageHandler)
	router.AddNoPublisherHandler("mpv2", "out", pubSub, func(msg *message.Message) error {
		expected := `{"client_id":"C123456","consent":{"ad_user_data":"GRANTED","ad_personalization":"GRANTED","analytics_storage":"GRANTED"},"events":[{"name":"add_to_cart","params":{"page_location":"https://foomo.org","page_title":"Home"}}],"debug_mode":true}`
		if !assert.JSONEq(t, expected, string(msg.Payload)) {
			fmt.Println(string(msg.Payload))
		}
		done.Store(true)
		return nil
	})

	go func() {
		assert.NoError(t, router.Run(context.TODO()))
	}()
	assert.Eventually(t, router.IsRunning, time.Second, 50*time.Millisecond)

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

	msg := message.NewMessage(watermill.NewUUID(), jsonPayload)

	require.NoError(t, pubSub.Publish("in", msg))

	assert.Eventually(t, done.Load, time.Second, 50*time.Millisecond)
}
