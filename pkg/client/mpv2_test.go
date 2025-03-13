package client_test

import (
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"testing"

	"github.com/foomo/sesamy-go/pkg/client"
	"github.com/foomo/sesamy-go/pkg/event"
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

func TestNewMPv2(t *testing.T) {
	l := zaptest.NewLogger(t)

	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Helper()
		out, err := httputil.DumpRequest(r, true)
		if assert.NoError(t, err) {
			t.Log(string(out))
		}
	}))

	c := client.NewMPv2(l, s.URL)
	incomingReq, err := http.NewRequestWithContext(t.Context(), http.MethodGet, "/foo/bar", nil)
	require.NoError(t, err)

	err = c.Collect(incomingReq, event.NewPageView(params.PageView{
		PageTitle:    "foo",
		PageLocation: "bar",
	}))
	require.NoError(t, err)
}
