package client_test

import (
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"testing"

	"github.com/foomo/sesamy-go/pkg/client"
	"github.com/foomo/sesamy-go/pkg/encoding/gtag"
	"github.com/foomo/sesamy-go/pkg/sesamy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

func TestNewGtag(t *testing.T) {
	t.Parallel()
	l := zaptest.NewLogger(t)

	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Helper()
		out, err := httputil.DumpRequest(r, true)
		if assert.NoError(t, err) {
			t.Log(string(out))
		}
	}))

	c := client.NewGTag(l, s.URL, "GA-XXXXXX")
	incomingReq, err := http.NewRequestWithContext(t.Context(), http.MethodGet, "/foo/bar", nil)
	require.NoError(t, err)

	err = c.Send(incomingReq, &gtag.Payload{
		EventName: gtag.Set(sesamy.EventName("page_view")),
	})
	require.NoError(t, err)
}
