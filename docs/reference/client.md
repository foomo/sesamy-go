# Client

`pkg/client` provides direct GA4 senders for server-to-GA traffic.

## `client.MPv2`

### Constructor

```go
func NewMPv2(l *zap.Logger, host string, opts ...MPv2Option) *MPv2
```

Defaults:

- `path = "/mp/collect"`
- `cookies = ["gtm_auth", "gtm_debug", "gtm_preview"]`
- `protocolVersion = "2"`
- `httpClient = http.DefaultClient`
- Always appends `MPv2MiddlewarClientID` to the middleware chain.

### Options

| Option                                      | Purpose                                  |
| ------------------------------------------- | ---------------------------------------- |
| `MPv2WithHTTPClient(*http.Client)`          | Override HTTP client.                    |
| `MPv2WithPath(string)`                      | Override upstream path.                  |
| `MPv2WithCookies(...string)`                | Cookie names forwarded upstream.         |
| `MPv2WithAPISecret(string)`                 | GA4 Measurement Protocol API secret.     |
| `MPv2WithMeasurementID(string)`             | `G-XXXXXXXXX`.                           |
| `MPv2WithMiddlewares(...MPv2Middleware)`    | Pre-send enrichment middleware.          |

### Methods

```go
func (c *MPv2) Collect(r *http.Request, events ...sesamy.AnyEvent) error
func (c *MPv2) SendPayload(r *http.Request, payload *mpv2.Payload[any]) error
func (c *MPv2) HTTPClient() *http.Client
```

`Collect` builds a `*mpv2.Payload[any]` from the events and runs `SendPayload`.

### Middleware types

```go
type MPv2Handler    func(r *http.Request, payload *mpv2.Payload[any]) error
type MPv2Middleware func(next MPv2Handler) MPv2Handler
```

Built-in: `MPv2MiddlewarClientID` (auto-fills client ID from `_ga` cookie).

---

## `client.GTag`

### Constructor

```go
func NewGTag(l *zap.Logger, host, trackingID string, opts ...GTagOption) *GTag
```

Defaults:

- `path = "/g/collect"`
- `cookies = ["gtm_auth", "gtm_debug", "gtm_preview"]`
- `protocolVersion = "2"`
- `httpClient = http.DefaultClient`
- Middleware chain auto-appends: `GTagMiddlewareRichsstsse`, `GTagMiddlewarTrackingID`, `GTagMiddlewarProtocolVersion`, `GTagMiddlewarIsDebug`, `GTagMiddlewarClientID`, `GTagMiddlewarSessionID`.

### Options

| Option                                      | Purpose                                  |
| ------------------------------------------- | ---------------------------------------- |
| `GTagWithHTTPClient(*http.Client)`          | Override HTTP client.                    |
| `GTagWithPath(string)`                      | Override upstream path.                  |
| `GTagWithCookies(...string)`                | Cookies forwarded upstream.              |
| `GTagWithMiddlewares(...GTagMiddleware)`    | Pre-send middleware.                     |

### Methods

```go
func (c *GTag) Send(r *http.Request, payload *gtag.Payload) error
func (c *GTag) SendRaw(r *http.Request, payload *gtag.Payload) error
func (c *GTag) HTTPClient() *http.Client
```

`Send` runs the middleware chain then `SendRaw`. `SendRaw` POSTs the encoded payload to `${host}${path}`.

### Middleware types

```go
type GTagHandler    func(r *http.Request, payload *gtag.Payload) error
type GTagMiddleware func(next GTagHandler) GTagHandler
```
