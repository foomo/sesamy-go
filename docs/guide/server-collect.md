# Server-side Collect

`pkg/collect` is the high-level entry point for running a server-side tagging endpoint. It accepts inbound GA4 traffic on both protocols, runs your middleware chain, and forwards enriched payloads to a downstream tagging URL.

## Constructor

```go
import "github.com/foomo/sesamy-go/pkg/collect"

c, err := collect.New(l,
	collect.WithTagging("https://sgtm.example.com"),
	collect.WithTaggingClient(myHTTPClient),
	collect.WithGTagHTTPMiddlewares(gtagMiddlewares...),
	collect.WithMPv2HTTPMiddlewares(mpv2Middlewares...),
)
```

Options:

| Option                     | Purpose                                                    |
| -------------------------- | ---------------------------------------------------------- |
| `WithTagging(url)`         | Downstream tagging server URL (sGTM, vendor endpoint, GA4).|
| `WithTaggingClient(c)`     | Custom `*http.Client` for the upstream call.               |
| `WithGTagHTTPMiddlewares`  | Middleware chain for `/g/collect` (gtag.js).               |
| `WithMPv2HTTPMiddlewares`  | Middleware chain for `/mp/collect` (MPv2).                 |

## Handlers

The `Collect` value exposes two `http.HandlerFunc`s:

```go
c.GTagHTTPHandler   // ResponseWriter, *Request
c.MPv2HTTPHandler   // ResponseWriter, *Request
```

Mount them on any router (`net/http`, `chi`, `gin`, `echo`, ...):

```go
mux.HandleFunc("/g/collect", c.GTagHTTPHandler)
mux.HandleFunc("/mp/collect", c.MPv2HTTPHandler)
```

## Request flow

```
client ─▶ /g/collect or /mp/collect
            │
            ▼
       Handler decodes payload
            │
            ▼
       middleware chain (your enrichment, validation, side-effects)
            │
            ▼
       gtagHandler / mpv2Handler  (re-encodes payload, POSTs to taggingURL)
            │
            ▼
       upstream tagging server (sGTM / GA4 / vendor)
```

The terminal forward handler:

- Clones inbound `r.Header` to the upstream request.
- Re-encodes the payload (GTag → form values, MPv2 → JSON body).
- POSTs to `${taggingURL}/g/collect` or `${taggingURL}/mp/collect`.

## Cross-protocol event hooks

If you want to react to events using **typed GA4 events** regardless of protocol, wrap an `EventHandler` with `MiddlewareEventHandler` from either subpackage:

```go
import (
	sesamyhttp "github.com/foomo/sesamy-go/pkg/http"
	gtaghttp "github.com/foomo/sesamy-go/pkg/http/gtag"
	mpv2http "github.com/foomo/sesamy-go/pkg/http/mpv2"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

handler := sesamyhttp.EventHandler(
	func(l *zap.Logger, r *http.Request, e *sesamy.Event[any]) error {
		l.Info("event", zap.String("name", e.Name.String()))
		return nil
	},
)

collect.New(l,
	collect.WithTagging("https://sgtm.example.com"),
	collect.WithGTagHTTPMiddlewares(gtaghttp.MiddlewareEventHandler(handler)),
	collect.WithMPv2HTTPMiddlewares(mpv2http.MiddlewareEventHandler(handler)),
)
```

Under the hood, the GTag variant converts incoming GTag → MPv2 before calling your handler, then re-encodes back to GTag for forwarding.

## Recommended middleware stack

For a typical sGTM-fronting setup:

**GTag (`/g/collect`):**

```go
gtaghttp.MiddlewareLogger,
gtaghttp.MiddlewareUserID("uid"),         // if you store a first-party user id cookie
gtaghttp.MiddlewareWithTimeout(2*time.Second),
```

**MPv2 (`/mp/collect`):**

```go
mpv2http.MiddlewareLogger,
mpv2http.MiddlewareClientID,              // populates from _ga cookie if missing
mpv2http.MiddlewareSessionID(measurementID),
mpv2http.MiddlewareIPOverride,            // CF-Connecting-IP / X-Forwarded-For / ...
mpv2http.MiddlewareUserAgent,
mpv2http.MiddlewarePageLocation,          // Referer / X-Page-Title / X-Page-Referrer
mpv2http.MiddlewareEngagementTime,        // default 100ms if missing
mpv2http.MiddlewareDebugMode,
mpv2http.MiddlewareWithTimeout(2*time.Second),
```

Full reference: [reference/http](/reference/http).
