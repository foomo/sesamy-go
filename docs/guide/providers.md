# Providers

Provider packages adapt third-party systems (consent, marketing, attribution) to the `sesamy-go` event model. Each lives under `pkg/provider/<name>`.

## Available providers

| Provider                                           | Purpose                                                   | Package                       |
| -------------------------------------------------- | --------------------------------------------------------- | ----------------------------- |
| [Cookiebot](/reference/providers/cookiebot)        | Parse the `CookieConsent` cookie into typed flags.        | `pkg/provider/cookiebot`      |
| [Emarsys](/reference/providers/emarsys)            | Marketing events: cart, category, purchase, view.         | `pkg/provider/emarsys`        |
| [Tracify](/reference/providers/tracify)            | Attribution / conversion tracking events.                 | `pkg/provider/tracify`        |

## How providers fit in

Two integration shapes:

### 1. Consent (Cookiebot)

Parses an inbound cookie into a struct you can branch on in middleware:

```go
import "github.com/foomo/sesamy-go/pkg/provider/cookiebot"

func consentMiddleware(next mpv2http.MiddlewareHandler) mpv2http.MiddlewareHandler {
	return func(l *zap.Logger, w http.ResponseWriter, r *http.Request, p *mpv2.Payload[any]) error {
		var c cookiebot.Cookie
		if raw, _ := r.Cookie(cookiebot.CookieName); raw != nil {
			_ = json.Unmarshal([]byte(raw.Value), &c)
		}
		if !c.Statistics {
			return nil // drop event
		}
		return next(l, w, r, p)
	}
}
```

### 2. Typed event constructors (Emarsys, Tracify)

Define typed `Event[P]` values that map onto vendor schemas:

```go
import (
	emarsysevent "github.com/foomo/sesamy-go/pkg/provider/emarsys/event"
	"github.com/foomo/sesamy-go/pkg/provider/emarsys/event/params"
)

ev := emarsysevent.NewEmarsysCart(params.EmarsysCart[sesamyparams.Item]{
	// ...
})
```

These integrate naturally with `client.MPv2.Collect` for direct server-to-vendor sends, or as branching targets inside a middleware-driven event handler.

## Adding a new provider

To wire a new vendor:

1. Create `pkg/provider/<name>/`.
2. Define event-name constants and typed params under `event/` and `event/params/`.
3. Provide constructors that return `sesamy.Event[P]`.
4. Optionally provide a parser (cookie, header) that returns a typed config.
5. Add a page under `docs/reference/providers/<name>.md` and link it from this guide + `docs/.vitepress/config.mts` sidebar.
