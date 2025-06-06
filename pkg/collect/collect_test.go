package collect_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/foomo/sesamy-go/pkg/collect"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

func TestNew(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		opts    []collect.Option
		wantErr bool
	}{
		{
			name:    "create with default options",
			opts:    nil,
			wantErr: false,
		},
		{
			name: "create with custom tagging URL",
			opts: []collect.Option{
				collect.WithTagging("https://example.com"),
			},
			wantErr: false,
		},
		{
			name: "create with custom HTTP client",
			opts: []collect.Option{
				collect.WithTaggingClient(http.DefaultClient),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			l := zaptest.NewLogger(t)
			c, err := collect.New(l, tt.opts...)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.NotNil(t, c)
		})
	}
}

func TestCollect_GTagHTTPHandler(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name            string
		query           string
		setupMockServer func(query string) *httptest.Server
		expectedStatus  int
	}{
		{
			name:  "successful request",
			query: "v=2&tid=G-C5FH0JEWES&gtm=45he5641h1v9208481823z89208852519za204zb9208852519&_p=1749197326339&_dbg=1&gcs=G111&gcd=13t3t3t2t5l1&npa=0&dma_cps=syphamo&dma=1&tag_exp=101509157~103116026~103200004~103233427~103351869~103351871~104653070~104653072~104661466~104661468~104698127~104698129&gdid=dMWZhNz&cid=1897509477.1748871081&ecid=1224659014&ul=en-us&sr=1728x1117&_fplc=0&ur=DE&uaa=arm&uab=64&uafvl=Chromium%3B136.0.7103.114%7CGoogle%2520Chrome%3B136.0.7103.114%7CNot.A%252FBrand%3B99.0.0.0&uamb=0&uam=&uap=macOS&uapv=15.5.0&uaw=0&are=1&frm=0&pscdl=noapi&_eu=AAAAAAQ&sst.rnd=602677853.1749197331&sst.etld=google.de&sst.gcsub=region1&sst.adr=1&sst.us_privacy=1YNY&sst.tft=1749197326339&sst.lpc=223090308&sst.navt=r&sst.ude=0&sst.sw_exp=1&_s=1&sid=1749197325&sct=4&seg=1&dl=https%3A%2F%2Fnode-b.stage.geschenkidee.ch%2Fsets-geschenkkoerbe&dr=https%3A%2F%2Fnode-b.stage.geschenkidee.ch%2Fgeschenkgutscheine-2&dt=Geschenksets%202025%20%E2%80%93%20Kreative%20Sets%20f%C3%BCr%20jeden%20Anlass%20%7C%20geschenkidee.ch&_tu=DA&en=page_view&epn.emarsys_page_view_id=1550965197&tfd=4880&richsstsse",
			setupMockServer: func(query string) *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					assert.Equal(t, "/g/collect", r.URL.Path)
					if !assert.Len(t, r.URL.RawQuery, len(query)) {
						t.Logf("expected: %s", query)
						t.Logf("actual:   %s", r.URL.RawQuery)
					}
					w.WriteHeader(http.StatusOK)
				}))
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:  "server error",
			query: "",
			setupMockServer: func(query string) *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusInternalServerError)
				}))
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			l := zaptest.NewLogger(t)

			// Setup mock server
			mockServer := tt.setupMockServer(tt.query)
			defer mockServer.Close()

			// Create collector with mock server URL
			c, err := collect.New(l, collect.WithTagging(mockServer.URL))
			require.NoError(t, err)

			// Create test request
			req := httptest.NewRequest(http.MethodPost, "/collect?"+tt.query, nil)
			req.Header.Set("User-Agent", "test-agent")

			// Create response recorder
			w := httptest.NewRecorder()

			// Call the handler
			c.GTagHTTPHandler(w, req)

			// Assert response
			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}

//
// func TestCollect_WithMiddleware(t *testing.T) {
// 	logger := zap.NewNop()
//
// 	// Create a test middleware
// 	testMiddleware := func(next gtaghttp.HandlerFunc) gtaghttp.HandlerFunc {
// 		return func(l *zap.Logger, w http.ResponseWriter, r *http.Request, payload *gtag.Payload) error {
// 			// Modify the payload or do something before passing to next handler
// 			payload.ClientID = "test-client"
// 			return next(l, w, r, payload)
// 		}
// 	}
//
// 	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		assert.Contains(t, r.URL.String(), "cid=test-client")
// 		w.WriteHeader(http.StatusOK)
// 	}))
// 	defer mockServer.Close()
//
// 	c, err := New(logger,
// 		WithTagging(mockServer.URL),
// 		WithGTagHTTPMiddlewares(testMiddleware),
// 	)
// 	require.NoError(t, err)
//
// 	// Create test request
// 	req := httptest.NewRequest(http.MethodPost, "/collect", nil)
// 	w := httptest.NewRecorder()
//
// 	// Call the handler
// 	c.GTagHTTPHandler(w, req)
//
// 	assert.Equal(t, http.StatusOK, w.Code)
// }
