# Core Concepts

## Events

The core type is generic:

```go
type Event[P any] struct {
	Name   EventName `json:"name"`
	Params P         `json:"params,omitempty"`
}
```

- `EventName` is a string type with 44 predefined GA4 constants (`EventNamePageView`, `EventNamePurchase`, ...) — see [reference/events](/reference/encoding#event-names).
- `P` is the typed parameter struct — provider events define their own (e.g. `params.EmarsysCart`).
- `Event[P]` can be projected to `Event[any]` (a.k.a. `AnyEvent`) via `.AnyEvent()` for protocol-level encoding.

```go
e := sesamy.NewEvent(sesamy.EventNamePurchase, MyPurchaseParams{...})
any := e.AnyEvent()   // for transport
```

## Protocols

GA4 traffic comes in two shapes; `sesamy-go` handles both.

| Protocol | Encoding              | Used by                                    | Package                         |
| -------- | --------------------- | ------------------------------------------ | ------------------------------- |
| GTag     | form-encoded querystring | `gtag.js` in the browser                | `pkg/encoding/gtag`             |
| MPv2     | JSON body             | Measurement Protocol senders, server SDKs  | `pkg/encoding/mpv2`             |

Conversion helpers:

- `gtagencode.MPv2` — GTag → MPv2
- `mpv2encode.GTag` — MPv2 → GTag

This lets middleware operate on MPv2 events (richer, typed) regardless of which protocol the request arrived on.

## Handlers

Per-protocol HTTP handlers live in `pkg/http/gtag` and `pkg/http/mpv2`. Each exposes:

```go
func Handler(w http.ResponseWriter, r *http.Request) *Payload
```

Returning `nil` if the request is invalid (it also writes an error response).

The shared cross-protocol handler type is in `pkg/http`:

```go
type EventHandler func(l *zap.Logger, r *http.Request, event *sesamy.Event[any]) error
```

`MiddlewareEventHandler(EventHandler)` wraps an `EventHandler` so it works against the per-protocol middleware chain.

## Middleware

Each protocol defines:

```go
type MiddlewareHandler func(l *zap.Logger, w http.ResponseWriter, r *http.Request, payload *Payload) error
type Middleware        func(next MiddlewareHandler) MiddlewareHandler
```

Composition is right-to-left (last middleware runs first). The terminal handler is provided by your sink (e.g. `Collect` forwards to a tagging URL).

## Collect

`pkg/collect` ties it all together: a `Collect` struct with `GTagHTTPHandler` / `MPv2HTTPHandler` methods, an output `taggingURL`, and per-protocol middleware lists. See [server-side collect](/guide/server-collect).

## Clients

`pkg/client` provides direct senders for server-to-GA traffic (no browser involved):

- `client.GTag` — sends GTag payloads to a `/g/collect` endpoint.
- `client.MPv2` — sends MPv2 payloads to a `/mp/collect` endpoint with API secret + measurement ID.

Both have their own middleware chains (for client-side enrichment before send).

## Providers

`pkg/provider` holds vendor-specific event types and parsers:

- **Cookiebot** — parses the `CookieConsent` cookie.
- **Emarsys** — typed events: `emarsys_cart`, `emarsys_category`, `emarsys_purchase`, `emarsys_view`.
- **Tracify** — conversion-tracking events.

## Sessions

`pkg/session` provides helpers to parse `_ga` (client ID) and `_ga_<measurement_id>` (session ID + number) cookies. Middleware uses these to populate fields on inbound events that arrive without them.
