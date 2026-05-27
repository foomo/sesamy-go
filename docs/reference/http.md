# HTTP handlers and middleware

## Shared types — `pkg/http`

```go
type EventHandler func(l *zap.Logger, r *http.Request, event *sesamy.Event[any]) error
```

Protocol-neutral handler. Use it via `MiddlewareEventHandler` in either protocol subpackage to react to events with typed GA4 semantics regardless of wire format.

---

## GTag — `pkg/http/gtag`

### `Handler`

```go
func Handler(w http.ResponseWriter, r *http.Request) *gtag.Payload
```

- Supports `GET` and `POST`.
- `POST` reads either querystring or form-encoded body.
- Returns `nil` and writes an HTTP error if the payload is invalid (missing `en` event name).

### Middleware types

```go
type MiddlewareHandler func(l *zap.Logger, w http.ResponseWriter, r *http.Request, payload *gtag.Payload) error
type Middleware        func(next MiddlewareHandler) MiddlewareHandler
```

### Provided middlewares

| Name                              | Effect                                                              |
| --------------------------------- | ------------------------------------------------------------------- |
| `MiddlewareEventHandler(h)`       | Converts payload to MPv2, runs the cross-protocol `EventHandler` per event, re-encodes back to GTag. |
| `MiddlewareUserID(cookieName)`    | Sets `payload.UserID` from a named cookie if present.               |
| `MiddlewareWithTimeout(d)`        | Wraps the downstream call in a timeout context.                     |
| `MiddlewareLogger`                | Adds `event_name`, `event_user_id`, `event_client_id`, `event_session_id` to the zap logger. |

---

## MPv2 — `pkg/http/mpv2`

### `Handler`

```go
func Handler(w http.ResponseWriter, r *http.Request) *mpv2.Payload[any]
```

- `POST` only.
- Body is JSON-decoded into `*mpv2.Payload[any]`.
- Returns `nil` and writes an HTTP error if there are no events or any event is missing `name`.

### Middleware types

```go
type MiddlewareHandler func(l *zap.Logger, w http.ResponseWriter, r *http.Request, payload *mpv2.Payload[any]) error
type Middleware        func(next MiddlewareHandler) MiddlewareHandler
```

### Provided middlewares

| Name                              | Effect                                                              |
| --------------------------------- | ------------------------------------------------------------------- |
| `MiddlewareEventHandler(h)`       | Runs the cross-protocol `EventHandler` per event in the payload.    |
| `MiddlewareClientID`              | Fills `payload.ClientID` from the `_ga` cookie if empty.            |
| `MiddlewareSessionID(measurementID)` | Fills `ga_session_id` / `ga_session_number` per event from `_ga_<id>` cookie. |
| `MiddlewareUserAgent`             | Sets `user_agent` event param from `User-Agent` header.             |
| `MiddlewareIPOverride`            | Sets `ip_override` from `CF-Connecting-IP` / `X-Original-Forwarded-For` / `X-Forwarded-For` / `X-Real-Ip`. |
| `MiddlewarePageLocation`          | Sets `page_location` from `Referer`, `page_title` from `X-Page-Title`, `page_referrer` from `X-Page-Referrer`. |
| `MiddlewareEngagementTime`        | Sets `engagement_time_msec=100` per event if absent.                |
| `MiddlewareDebugMode`             | Preserves the debug flag.                                           |
| `MiddlewareWithTimeout(d)`        | Wraps downstream call in a timeout context (uses `context.WithoutCancel(r.Context())` so upstream cancellation doesn't kill it). |
| `MiddlewareLogger`                | Adds event names + OpenTelemetry `trace_id` / `span_id` to logs if a span is active. |

## Composition

Middleware compose with right-to-left order; the chain terminates at a "sink" handler (e.g. `Collect`'s `gtagHandler` / `mpv2Handler`, which forwards to the configured tagging URL).

```go
next := finalHandler
for _, m := range middlewares {
	next = m(next)
}
_ = next(l, w, r, payload)
```
