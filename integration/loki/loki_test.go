package loki_test

import (
	"testing"
	"time"

	"github.com/foomo/sesamy-go/integration/loki"
	"github.com/foomo/sesamy-go/pkg/encoding/mpv2"
	"github.com/foomo/sesamy-go/pkg/sesamy"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

func TestLoki_WriteDropsWhenBufferFull(t *testing.T) {
	t.Parallel()

	core, logs := observer.New(zap.WarnLevel)
	// The Loki is not started, so nothing reads the buffer of one line.
	l := loki.New(zap.New(core), "http://loki.invalid:3100", loki.WithBufferSize(1))

	payload := mpv2.Payload[any]{
		ClientID: "client",
		Events: []sesamy.Event[any]{
			{Name: "page_view"},
			{Name: "click"},
			{Name: "scroll"},
		},
	}

	done := make(chan struct{})

	go func() {
		l.Write(payload)
		l.Write(payload)
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("Write blocked on a full buffer")
	}

	assert.Equal(t, uint64(5), l.Dropped())
	assert.Equal(t, 1, logs.FilterMessage("buffer full, dropping lines").Len(), "warns once per interval")
}
