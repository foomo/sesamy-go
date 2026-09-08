# Loki integration

`integration/loki` ships server-side events to [Grafana Loki](https://grafana.com/oss/loki/) via the Loki push API (protobuf + snappy compression). Useful for retaining a high-volume, queryable log of every event flowing through your collect endpoint.

## Features

- **Protobuf + snappy** wire format — the efficient path Loki supports.
- **Channel-based batching** — events are queued and shipped in batches.
- **Exponential backoff** retry on transport errors.
- **Lifecycle service** — start/stop hooks for graceful shutdown (`Service.Start` / `Service.Close`).

## Wire it as a service

```go
import (
	"github.com/foomo/sesamy-go/integration/loki"
)

l, _ := zap.NewProduction()

lk := loki.New(l, loki.WithEndpoint("https://logs.example.com/loki/api/v1/push"))
svc := loki.NewService(lk)

// In your lifecycle manager:
go func() { _ = svc.Start(ctx) }()
defer svc.Close(ctx)
```

## Plug into Collect

Wrap a Loki push in a middleware so every successfully forwarded event lands in Loki:

```go
import (
	lokimw "github.com/foomo/sesamy-go/integration/loki"
)

collect.New(l,
	collect.WithTagging("https://sgtm.example.com"),
	collect.WithMPv2HTTPMiddlewares(
		lokimw.MPv2Middleware(lk), // see integration/loki/middleware.go
		mpv2http.MiddlewareLogger,
	),
)
```

## Labels and lines

Each event becomes a Loki **line** with metadata labels (event name, measurement ID, client ID where available). The line format is defined in `integration/loki/line.go` and can be customized by replacing the line builder.

## Operational notes

- Backpressure: the channel has a bounded buffer; if Loki is unreachable for too long, oldest events are dropped to keep the producer non-blocking. Tune the buffer size to your retention tolerance.
- Compression: snappy is mandatory for the Loki protobuf endpoint; do not disable it.
- Multi-tenancy: pass `X-Scope-OrgID` via a custom HTTP client / request options if your Loki is multi-tenant.
