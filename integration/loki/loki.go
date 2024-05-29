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

// Loki is a io.Writer, that writes given log entries by pushing
// directly to the given loki server URL. Each `Push` instance handles for a single tenant.
// No batching of log lines happens when sending to Loki.
type (
	Loki struct {
		l        *zap.Logger
		endpoint string

		// channel for incoming logs
		entries    chan logproto.Entry
		batchSize  int
		bufferSize int

		// shutdown channels
		cancel context.CancelFunc

		userAgent  string
		httpClient *http.Client
		backoff    *backoff.Config
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
func (p *Loki) Start(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	p.cancel = cancel
	utils.Batch(ctx, p.entries, p.batchSize, p.process)
}

func (p *Loki) Write(ts time.Time, payload mpv2.Payload[any]) {
	var metadata push.LabelsAdapter
	if payload.UserID != "" {
		metadata = append(metadata, push.LabelAdapter{
			Name:  "user_id",
			Value: payload.UserID,
		})
	}
	for _, event := range payload.Events {
		metadata := append(metadata, push.LabelAdapter{
			Name:  "name",
			Value: event.Name.String(),
		})

		line := Line{
			Params:             event.Params,
			ClientID:           payload.ClientID,
			UserID:             payload.UserID,
			UserProperties:     payload.UserProperties,
			Consent:            payload.Consent,
			NonPersonalizedAds: payload.NonPersonalizedAds,
			UserData:           payload.UserData,
			DebugMode:          payload.DebugMode,
		}

		lineBytes, err := line.Marshal()
		if err != nil {
			p.l.Warn("failed to marshal line", zap.Error(err))
			continue
		}

		p.entries <- logproto.Entry{
			Line:               string(lineBytes),
			Timestamp:          time.UnixMicro(payload.TimestampMicros),
			StructuredMetadata: metadata,
		}
	}
}

// Stop will cancel any ongoing requests and stop the goroutine listening for requests
func (p *Loki) Stop() {
	if p.cancel != nil {
		p.cancel()
		p.cancel = nil
	}
}

// ------------------------------------------------------------------------------------------------
// ~ Private methods
// ------------------------------------------------------------------------------------------------

func (p *Loki) process(entries []logproto.Entry) {
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
		p.l.Error("failed to marshal payload to json", zap.Error(err))
		return
	}

	payload := snappy.Encode(nil, request)

	// We will use a timeout within each attempt to send
	back := backoff.New(context.Background(), *p.backoff)

	// send log with retry
	for {
		var status int
		status, err = p.send(context.Background(), payload)
		if err == nil {
			break
		}

		if status > 0 && status != 429 && status/100 != 5 {
			p.l.Error("failed to send entry, server rejected push with a non-retryable status code", zap.Error(err), zap.Int("status", status))
			break
		}

		if !back.Ongoing() {
			p.l.Error("failed to send entry, retries exhausted, entry will be dropped", zap.Error(err), zap.Int("status", status))
			break
		}
		p.l.Warn("failed to send entry, retrying", zap.Error(err), zap.Int("status", status))
		back.Wait()
	}
}

// send makes one attempt to send the payload to Loki
func (p *Loki) send(ctx context.Context, payload []byte) (int, error) {
	// Set a timeout for the request
	ctx, cancel := context.WithTimeout(ctx, p.httpClient.Timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, p.endpoint, bytes.NewReader(payload))
	if err != nil {
		return -1, errors.Wrap(err, "failed to create request")
	}
	req.Header.Set("Content-Type", defaultContentType)
	req.Header.Set("User-Agent", p.userAgent)

	resp, err := p.httpClient.Do(req)
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
		p.l.Error("failed to close response body", zap.Error(err))
	}

	return status, err
}
