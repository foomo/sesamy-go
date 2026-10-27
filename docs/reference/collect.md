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
WithMeasurementID(id string)                             // default MPv2 measurement_id
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

1. Decode payload using the per-protocol handler — invalid requests get an HTTP error and never reach middleware. GTag uses `HandlerBatch`, so a batched `POST` body (one event per line) yields one payload per line.
2. Run the configured middleware chain once per payload.
3. Re-encode the payload and `POST` it to `${taggingURL}${path}`, where path is `/g/collect` for GTag and `/mp/collect` for MPv2. Inbound `r.Header` is cloned to the upstream request. For MPv2 the inbound query is kept; if it has no `measurement_id`, the value from `WithMeasurementID` is set.
4. Copy the upstream status code, headers and body to the inbound response. For batched GTag requests only the first payload's upstream response is written.
5. Errors from middleware or the upstream call map to `500 Internal Server Error`.

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
