# Sending events from Go

`pkg/client` lets a Go process emit GA4 events directly, without going through a browser. Useful for backend-only conversions (refunds, async order completion, webhook-driven events).

Two clients, one per protocol:

- `client.GTag` — sends form-encoded payloads to a `/g/collect` endpoint.
- `client.MPv2` — sends JSON to a `/mp/collect` endpoint with API secret + measurement ID.

## MPv2 client (recommended for server-side)

```go
import (
	"github.com/foomo/sesamy-go/pkg/client"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

c := client.NewMPv2(l, "https://www.google-analytics.com",
	client.MPv2WithAPISecret("YOUR_SECRET"),
	client.MPv2WithMeasurementID("G-XXXXXXXXX"),
)

err := c.Collect(r,
	sesamy.NewEvent(sesamy.EventNamePurchase, map[string]any{
		"transaction_id": "T-123",
		"value":          42.0,
		"currency":       "EUR",
	}).AnyEvent(),
)
```

### Options

| Option                                  | Purpose                                              |
| --------------------------------------- | ---------------------------------------------------- |
| `MPv2WithHTTPClient(*http.Client)`      | Override the HTTP client (timeouts, transports).     |
| `MPv2WithPath(string)`                  | Override the upstream path (default `/mp/collect`).  |
| `MPv2WithCookies(...string)`            | Cookie names to forward from the inbound request.   |
| `MPv2WithAPISecret(string)`             | GA4 Measurement Protocol API secret.                 |
| `MPv2WithMeasurementID(string)`         | `G-XXXXXXXXX` from GA4 Admin → Data Streams.         |
| `MPv2WithMiddlewares(...MPv2Middleware)`| Pre-send enrichment middleware.                      |

### Methods

```go
c.Collect(r, events...)              // wraps events into a Payload and sends
c.SendPayload(r, payload)            // sends a fully built mpv2.Payload[any]
c.HTTPClient()                        // exposes underlying *http.Client
```

The constructor implicitly appends `MPv2MiddlewarClientID` so missing client IDs are sourced from the inbound `_ga` cookie.

## GTag client

```go
c := client.NewGTag(l, "https://www.google-analytics.com", "G-XXXXXXXXX",
	client.GTagWithCookies("gtm_auth"),
)

err := c.Send(r, &gtag.Payload{...})
```

### Options

| Option                                   | Purpose                                                          |
| ---------------------------------------- | ---------------------------------------------------------------- |
| `GTagWithHTTPClient(*http.Client)`       | Override HTTP client.                                            |
| `GTagWithPath(string)`                   | Override upstream path (default `/g/collect`).                   |
| `GTagWithCookies(...string)`             | Cookie names to forward.                                         |
| `GTagWithMiddlewares(...GTagMiddleware)` | Pre-send middleware.                                             |

Built-in middlewares appended in the constructor:

- `GTagMiddlewareRichsstsse` — populates the `richsstsse` parameter.
- `GTagMiddlewarTrackingID` — fills tracking ID if missing.
- `GTagMiddlewarProtocolVersion` — sets `v=2`.
- `GTagMiddlewarIsDebug` — preserves debug flag.
- `GTagMiddlewarClientID` — fills client ID from `_ga`.
- `GTagMiddlewarSessionID` — fills session ID from `_ga_<id>`.

## Choosing between them

- Use **MPv2** for new server-side code — it's the typed JSON protocol GA4 recommends and the easiest to compose middleware against.
- Use **GTag** only if you must mimic exact `gtag.js` behavior (e.g. when fronting a tagging server that expects gtag traffic).
