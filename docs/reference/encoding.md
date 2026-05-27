# Encoding

GA4 traffic comes in two protocols and `sesamy-go` round-trips between them.

## GTag (`pkg/encoding/gtag`)

Form-encoded querystring protocol used by `gtag.js`. Each field maps to a short key (e.g. `tid`, `cid`, `en`, `dl`).

```go
import "github.com/foomo/sesamy-go/pkg/encoding/gtag"

var p gtag.Payload
if err := gtag.Decode(url.Values{...}, &p); err != nil { /* ... */ }

values, body, err := gtag.Encode(&p)  // re-encode to url.Values + body
qs := gtag.EncodeValues(values)
```

### `Payload` struct (selected fields)

| Field              | GTag key | Description                                     |
| ------------------ | -------- | ----------------------------------------------- |
| `ProtocolVersion`  | `v`      | Protocol version (always `2`).                  |
| `TrackingID`       | `tid`    | Measurement / stream ID (`G-XXXXXXXXX`).        |
| `ClientID`         | `cid`    | GA client ID (`281344611.1635634925`).          |
| `UserID`           | `uid`    | Custom user ID.                                 |
| `SessionID`        | `sid`    | Session ID.                                     |
| `EventName`        | `en`     | Event name (e.g. `page_view`).                  |
| `DocumentLocation` | `dl`     | Current page URL.                               |
| `DocumentTitle`    | `dt`     | Current page title.                             |
| `DocumentReferrer` | `dr`     | Referrer.                                       |
| `IsDebug`          | `_dbg`   | Debug flag.                                     |
| `GTMHashInfo`      | `gtm`    | GTM container fingerprint.                      |

Embedded sub-structs (with `,inline,squash`):

- `Consent` — consent-mode signals.
- `Campaign` — UTM-style campaign fields.
- `ECommerce` — items, value, currency.
- `ClientHints` — UA-CH derived.

Full field list lives in `pkg/encoding/gtag/payload.go`.

### Helpers

```go
gtag.Get[T any](*T, T) T              // safe deref with fallback
gtag.GetDefault[T any](*T, T) T       // alias
```

## MPv2 (`pkg/encoding/mpv2`)

JSON body protocol used by GA4's Measurement Protocol v2.

```go
import "github.com/foomo/sesamy-go/pkg/encoding/mpv2"

p := mpv2.NewPayload[any]()
p.Events = append(p.Events, sesamy.NewEvent(sesamy.EventNamePageView, params).AnyEvent())
body, _ := json.Marshal(p)
```

### `Payload[P]` (selected fields)

| Field            | JSON key            | Description                                  |
| ---------------- | ------------------- | -------------------------------------------- |
| `ClientID`       | `client_id`         | GA client ID.                                |
| `UserID`         | `user_id`           | Custom user ID.                              |
| `SessionID`      | (per event param)   | Session ID lives on each event's params.     |
| `TimestampMicros`| `timestamp_micros`  | Hit timestamp in microseconds.               |
| `Events`         | `events`            | Slice of `sesamy.Event[P]`.                  |
| `UserProperties` | `user_properties`   | User-scoped properties.                      |
| `Consent`        | `consent`           | Consent flags.                               |
| `UserData`       | `user_data`         | Enhanced conversions (SHA-256 hashed).       |

Also: `userdata.go`, `userdataaddress.go`, `sha256hash.go` — helpers for **enhanced conversions** (server-side hashing of email/phone/address).

## Conversion

### GTag → MPv2

```go
import "github.com/foomo/sesamy-go/pkg/encoding/gtagencode"

var p *mpv2.Payload[any]
err := gtagencode.MPv2(gtagPayload, &p)
```

### MPv2 → GTag

```go
import "github.com/foomo/sesamy-go/pkg/encoding/mpv2encode"

var p *gtag.Payload
err := mpv2encode.GTag[any](mpv2Payload, &p)
```

Used internally by `gtaghttp.MiddlewareEventHandler` so cross-protocol handlers can write MPv2-shaped logic regardless of inbound wire format.
