package gtag_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/foomo/sesamy-go/pkg/encoding/gtag"
	gtaghttp "github.com/foomo/sesamy-go/pkg/http/gtag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

func TestMiddlewareUserID_EmptyCookie(t *testing.T) {
	t.Parallel()

	r := httptest.NewRequestWithContext(t.Context(), http.MethodPost, "/", nil)
	r.AddCookie(&http.Cookie{Name: "uid", Value: ""})

	payload := &gtag.Payload{}
	next := func(*zap.Logger, http.ResponseWriter, *http.Request, *gtag.Payload) error { return nil }

	require.NoError(t, gtaghttp.MiddlewareUserID("uid")(next)(zaptest.NewLogger(t), httptest.NewRecorder(), r, payload))
	assert.Nil(t, payload.UserID)
}
