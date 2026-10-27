package client_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/foomo/sesamy-go/pkg/client"
	"github.com/foomo/sesamy-go/pkg/encoding/gtag"
	"github.com/foomo/sesamy-go/pkg/encoding/mpv2"
	"github.com/foomo/sesamy-go/pkg/sesamy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

func TestSendRaw_StripsCookiesAndAuthorization(t *testing.T) {
	t.Parallel()

	var got http.Header

	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got = r.Header.Clone()
	}))
	defer s.Close()

	newRequest := func() *http.Request {
		r := httptest.NewRequestWithContext(t.Context(), http.MethodGet, "/", nil)
		r.Header.Set("Cookie", "session=secret; gtm_debug=1")
		r.Header.Set("Authorization", "Bearer secret")

		return r
	}

	assertHeader := func() {
		assert.Equal(t, []string{"gtm_debug=1"}, got.Values("Cookie"))
		assert.Empty(t, got.Get("Authorization"))
	}

	l := zaptest.NewLogger(t)

	require.NoError(t, client.NewGTag(l, s.URL, "G-XXXXXX").SendRaw(newRequest(), &gtag.Payload{EventName: new(sesamy.EventNamePageView)}))
	assertHeader()

	require.NoError(t, client.NewMPv2(l, s.URL).SendRaw(newRequest(), mpv2.NewPayload[any]()))
	assertHeader()
}
