package mpv2_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/foomo/sesamy-go/pkg/encoding/mpv2"
	mpv2http "github.com/foomo/sesamy-go/pkg/http/mpv2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

func TestMiddlewareUserID_NoCookie(t *testing.T) {
	t.Parallel()

	r := httptest.NewRequestWithContext(t.Context(), http.MethodPost, "/", nil)
	payload := mpv2.NewPayload[any]()
	next := func(*zap.Logger, http.ResponseWriter, *http.Request, *mpv2.Payload[any]) error { return nil }

	require.NotPanics(t, func() {
		require.NoError(t, mpv2http.MiddlewareUserID("uid")(next)(zaptest.NewLogger(t), httptest.NewRecorder(), r, payload))
	})
	assert.Empty(t, payload.UserID)
}
