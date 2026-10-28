package gtag_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	gtaghttp "github.com/foomo/sesamy-go/pkg/http/gtag"
	"github.com/foomo/sesamy-go/pkg/sesamy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandlerBatch(t *testing.T) {
	t.Parallel()

	body := "en=page_view&_ss=1&_et=1\nen=scroll&epn.percent_scrolled=90&_et=2\n"
	r := httptest.NewRequestWithContext(t.Context(), http.MethodPost, "/g/collect?v=2&tid=G-TEST&cid=1.2&sid=3", strings.NewReader(body))
	w := httptest.NewRecorder()

	payloads := gtaghttp.HandlerBatch(w, r)
	require.Len(t, payloads, 2, w.Body.String())

	assert.Equal(t, sesamy.EventNamePageView, *payloads[0].EventName)
	assert.NotContains(t, payloads[0].EventParameterNumber, "percent_scrolled")
	assert.Equal(t, sesamy.EventNameScroll, *payloads[1].EventName)
	assert.Equal(t, "90", payloads[1].EventParameterNumber["percent_scrolled"])

	for _, p := range payloads {
		assert.Equal(t, "1.2", *p.ClientID)
		assert.Equal(t, "3", *p.SessionID)
	}
}

func TestHandler_SingleEvent(t *testing.T) {
	t.Parallel()

	r := httptest.NewRequestWithContext(t.Context(), http.MethodPost, "/g/collect?v=2&cid=1.2&en=page_view", nil)

	p := gtaghttp.Handler(httptest.NewRecorder(), r)
	require.NotNil(t, p)
	assert.Equal(t, sesamy.EventNamePageView, *p.EventName)
}
