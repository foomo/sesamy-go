# Collect

`pkg/collect`. High-level server that fronts GA4 traffic and forwards it to a tagging server.

## Constructor

```go
func New(l *zap.Logger, opts ...Option) (*Collect, error)
```

## Options

```go
WithTagging(url string)                                  // downstream tagging server URL
WithTaggingClient(c *http.Client)                        // override HTTP client
WithGTagHTTPMiddlewares(v ...gtaghttp.Middleware)        // /g/collect middleware chain
WithMPv2HTTPMiddlewares(v ...mpv2http.Middleware)        // /mp/collect middleware chain
```

## Methods

```go
func (c *Collect) GTagHTTPHandler(w http.ResponseWriter, r *http.Request)
func (c *Collect) MPv2HTTPHandler(w http.ResponseWriter, r *http.Request)
```

Both have signature `http.HandlerFunc` so they mount on any router.

## Behavior

For each request:

1. Decode payload using the per-protocol `Handler(...)` — invalid requests get an HTTP error and never reach middleware.
2. Run the configured middleware chain.
3. Re-encode the payload and `POST` it to `${taggingURL}${path}`, where path is `/g/collect` for GTag and `/mp/collect` for MPv2. Inbound `r.Header` is cloned to the upstream request.
4. Map upstream non-200 responses to `500 Internal Server Error` on the inbound side.

## Defaults

- `taggingClient` defaults to `http.DefaultClient` (override with `WithTaggingClient`).
- No middlewares by default. The chain is purely opt-in via the `With…Middlewares` options.

## Example

```go
c, err := collect.New(l,
	collect.WithTagging("https://sgtm.example.com"),
	collect.WithMPv2HTTPMiddlewares(
		mpv2http.MiddlewareLogger,
		mpv2http.MiddlewareClientID,
		mpv2http.MiddlewareIPOverride,
		mpv2http.MiddlewareUserAgent,
		mpv2http.MiddlewarePageLocation,
		mpv2http.MiddlewareEngagementTime,
		mpv2http.MiddlewareWithTimeout(2 * time.Second),
	),
)

http.HandleFunc("/g/collect", c.GTagHTTPHandler)
http.HandleFunc("/mp/collect", c.MPv2HTTPHandler)
```
