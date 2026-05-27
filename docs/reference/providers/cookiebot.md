# Cookiebot

`pkg/provider/cookiebot`.

Parses the [Cookiebot](https://www.cookiebot.com/) `CookieConsent` cookie into a typed Go struct.

## Cookie

```go
const CookieName = "CookieConsent"

type Cookie struct {
	Stamp       string `json:"stamp"`
	Necessary   bool   `json:"necessary"`
	Preferences bool   `json:"preferences"`
	Statistics  bool   `json:"statistics"`
	Marketing   bool   `json:"marketing"`
	Method      string `json:"method"`
	Version     string `json:"ver"`
	UTC         int    `json:"utc"`
	Region      string `json:"region"`
}
```

The raw cookie value looks like:

```
{stamp:'...',necessary:true,preferences:true,statistics:true,marketing:true,
 method:'explicit',ver:1,utc:1724770548958,region:'de'}
```

## Usage in middleware

Drop events when the visitor has not consented to statistics:

```go
import (
	"encoding/json"
	"net/http"

	"github.com/foomo/sesamy-go/pkg/provider/cookiebot"
)

func consentGate(next mpv2http.MiddlewareHandler) mpv2http.MiddlewareHandler {
	return func(l *zap.Logger, w http.ResponseWriter, r *http.Request, p *mpv2.Payload[any]) error {
		raw, err := r.Cookie(cookiebot.CookieName)
		if err != nil {
			return nil // no consent cookie → no event
		}

		var c cookiebot.Cookie
		if err := json.Unmarshal([]byte(raw.Value), &c); err != nil {
			return err
		}

		if !c.Statistics {
			return nil // statistics not allowed → drop
		}

		return next(l, w, r, p)
	}
}
```

## Subpackage

`pkg/provider/cookiebot/client` — HTTP client helpers for talking to Cookiebot's API (e.g. fetching consent reports). Refer to the source for the current surface.
