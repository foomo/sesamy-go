package loki

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/foomo/sesamy-go/pkg/encoding/mpv2"
	"github.com/foomo/sesamy-go/pkg/utils"
	"github.com/gogo/protobuf/proto"
	"github.com/golang/snappy"
	"github.com/grafana/dskit/backoff"
	"github.com/grafana/loki/pkg/push"
	"github.com/grafana/loki/v3/pkg/logproto"
	"github.com/pkg/errors"
	"github.com/prometheus/common/model"
	"go.uber.org/zap"
)

const (
	pushEndpoint               = "/loki/api/v1/push"
	defaultContentType         = "application/x-protobuf"
	defaultMaxReponseBufferLen = 1024
)

type (
	Loki struct {
		l        *zap.Logger
		endpoint string

		// channel for incoming logs
		entries chan logproto.Entry

		// shutdown
		cancel context.CancelFunc

		// options
		backoff    *backoff.Config
		batchSize  int
		bufferSize int
		userAgent  string
		httpClient *http.Client
	}
	Option func(*Loki)
)

// ------------------------------------------------------------------------------------------------
// ~ Options
// ------------------------------------------------------------------------------------------------

func WithHTTPClient(v *http.Client) Option {
	return func(l *Loki) {
		l.httpClient = v
	}
}

func WithBackoffConfig(v *backoff.Config) Option {
	return func(l *Loki) {
		l.backoff = v
	}
}

func WithUserAgent(v string) Option {
	return func(l *Loki) {
		l.userAgent = v
	}
}

func WithBatchSize(v int) Option {
	return func(l *Loki) {
		l.batchSize = v
	}
}

func WithBufferSize(v int) Option {
	return func(l *Loki) {
		l.bufferSize = v
	}
}

// ------------------------------------------------------------------------------------------------
// ~ Constructor
// ------------------------------------------------------------------------------------------------

// New creates an instance of `Push` which writes logs directly to given `lokiAddr`
func New(l *zap.Logger, addr string, opts ...Option) *Loki {
	inst := &Loki{
		l:          l,
		endpoint:   addr + pushEndpoint,
		httpClient: http.DefaultClient,
		userAgent:  "sesamy",
		batchSize:  10,
		bufferSize: 50,
		backoff: &backoff.Config{
			MinBackoff: 500 * time.Millisecond,
			MaxBackoff: 5 * time.Minute,
			MaxRetries: 10,
		},
	}

	for _, opt := range opts {
		if opt != nil {
			opt(inst)
		}
	}

	inst.entries = make(chan logproto.Entry, inst.bufferSize) // Use a buffered channel so we can retry failed pushes without blocking WriteEntry

	return inst
}

// Start pulls lines out of the channel and sends them to Loki
func (l *Loki) Start(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	l.cancel = cancel
	utils.Batch(ctx, l.entries, l.batchSize, l.process)
}

func (l *Loki) Write(payload mpv2.Payload[any]) {
	// sanity check
	if payload.ClientID == "" {
		l.l.Warn("received event without client id")
		return
	}

	for _, event := range payload.Events {
		// sanity check
		if event.Name == "" {
			l.l.Warn("received event without event name")
			continue
		}

		line := Line{
			Name:               event.Name,
			Params:             event.Params,
			UserID:             payload.UserID,
			Consent:            payload.Consent,
			UserData:           payload.UserData,
			ClientID:           payload.ClientID,
			UserProperties:     payload.UserProperties,
			NonPersonalizedAds: payload.NonPersonalizedAds,
			DebugMode:          payload.DebugMode,
		}

		lineBytes, err := line.Marshal()
		if err != nil {
			l.l.Warn("failed to marshal line", zap.Error(err))
			continue
		}

		if len(l.entries) == l.bufferSize {
			l.l.Warn("buffer size reached", zap.Int("size", l.bufferSize))
		}

		l.entries <- logproto.Entry{
			Line:      string(lineBytes),
			Timestamp: time.UnixMicro(payload.TimestampMicros),
			StructuredMetadata: push.LabelsAdapter{
				{
					Name:  "event_name",
					Value: event.Name.String(),
				},
			},
		}
	}
}

// Stop will cancel any ongoing requests and stop the goroutine listening for requests
func (l *Loki) Stop() {
	if l.cancel != nil {
		l.cancel()
		l.cancel = nil
	}
}

// ------------------------------------------------------------------------------------------------
// ~ Private methods
// ------------------------------------------------------------------------------------------------

func (l *Loki) process(entries []logproto.Entry) {
	l.l.Debug("processing entries batch", zap.Int("num", len(entries)))

	labels := model.LabelSet{
		"name":   "events",
		"stream": "sesamy",
	}
	request, err := proto.Marshal(&logproto.PushRequest{
		Streams: []logproto.Stream{
			{
				Labels:  labels.String(),
				Entries: entries,
				Hash:    uint64(labels.Fingerprint()),
			},
		},
	})
	if err != nil {
		l.l.Error("failed to marshal payload to json", zap.Error(err))
		return
	}

	payload := snappy.Encode(nil, request)

	// We will use a timeout within each attempt to send
	back := backoff.New(context.Background(), *l.backoff)

	// send log with retry
	for {
		var status int
		status, err = l.send(context.Background(), payload)
		if err == nil {
			break
		}

		if status > 0 && status != 429 && status/100 != 5 {
			l.l.Error("failed to send entries, server rejected push with a non-retryable status code", zap.Error(err), zap.Int("status", status))
			break
		}

		if !back.Ongoing() {
			l.l.Error("failed to send entries, retries exhausted, entries will be dropped", zap.Error(err), zap.Int("status", status))
			break
		}
		l.l.Warn("failed to send entries, retrying", zap.Error(err), zap.Int("status", status))
		back.Wait()
	}
}

// send makes one attempt to send the payload to Loki
func (l *Loki) send(ctx context.Context, payload []byte) (int, error) {
	// Set a timeout for the request
	ctx, cancel := context.WithTimeout(ctx, l.httpClient.Timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, l.endpoint, bytes.NewReader(payload))
	if err != nil {
		return -1, errors.Wrap(err, "failed to create request")
	}
	req.Header.Set("Content-Type", defaultContentType)
	req.Header.Set("User-Agent", l.userAgent)

	resp, err := l.httpClient.Do(req)
	if err != nil {
		return -1, errors.Wrap(err, "failed to send payload")
	}
	status := resp.StatusCode
	if status/100 != 2 {
		scanner := bufio.NewScanner(io.LimitReader(resp.Body, defaultMaxReponseBufferLen))
		line := ""
		if scanner.Scan() {
			line = scanner.Text()
		}
		err = fmt.Errorf("server returned HTTP status %s (%d): %s", resp.Status, status, line)
	}

	if err := resp.Body.Close(); err != nil {
		l.l.Error("failed to close response body", zap.Error(err))
	}

	return status, err
}
