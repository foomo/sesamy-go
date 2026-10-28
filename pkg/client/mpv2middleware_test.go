package client_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/foomo/sesamy-go/pkg/client"
	"github.com/foomo/sesamy-go/pkg/encoding/mpv2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMPv2MiddlewareUserID_NoCookie(t *testing.T) {
	t.Parallel()

	r := httptest.NewRequestWithContext(t.Context(), http.MethodGet, "/", nil)
	payload := mpv2.NewPayload[any]()
	next := func(*http.Request, *mpv2.Payload[any]) error { return nil }

	require.NotPanics(t, func() {
		require.NoError(t, client.MPv2MiddlewareUserID("uid")(next)(r, payload))
	})
	assert.Empty(t, payload.UserID)
}

func TestMPv2MiddlewarSessionID_WithoutPrefix(t *testing.T) {
	t.Parallel()

	require.NotPanics(t, func() {
		client.MPv2MiddlewarSessionID("F9XM71K45T")
	})
}
