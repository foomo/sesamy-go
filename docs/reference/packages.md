# Package overview

Go import paths and what each package contains.

| Import path                                          | Purpose                                                                                |
| ---------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `github.com/foomo/sesamy-go/pkg/sesamy`              | Core `Event[P]`, `EventName` constants, decode helpers.                                |
| `github.com/foomo/sesamy-go/pkg/event`               | Shared parameter types (Item, etc.) reused across providers.                           |
| `github.com/foomo/sesamy-go/pkg/encoding/gtag`       | GTag protocol encode/decode (form-encoded querystring).                                |
| `github.com/foomo/sesamy-go/pkg/encoding/mpv2`       | Measurement Protocol v2 payload types (JSON).                                          |
| `github.com/foomo/sesamy-go/pkg/encoding/gtagencode` | Convert GTag → MPv2.                                                                   |
| `github.com/foomo/sesamy-go/pkg/encoding/mpv2encode` | Convert MPv2 → GTag.                                                                   |
| `github.com/foomo/sesamy-go/pkg/http`                | Cross-protocol `EventHandler` type.                                                    |
| `github.com/foomo/sesamy-go/pkg/http/gtag`           | GTag HTTP handler + middleware chain.                                                  |
| `github.com/foomo/sesamy-go/pkg/http/mpv2`           | MPv2 HTTP handler + middleware chain.                                                  |
| `github.com/foomo/sesamy-go/pkg/client`              | Direct GTag/MPv2 clients for server-to-GA sends.                                       |
| `github.com/foomo/sesamy-go/pkg/collect`             | High-level `Collect` server: handlers + forwarding to a tagging server.                |
| `github.com/foomo/sesamy-go/pkg/session`             | Parse `_ga` (client ID), `_ga_<id>` (session ID + number).                             |
| `github.com/foomo/sesamy-go/pkg/utils`               | Miscellaneous helpers.                                                                 |
| `github.com/foomo/sesamy-go/pkg/provider/cookiebot`  | Cookiebot consent cookie type.                                                         |
| `github.com/foomo/sesamy-go/pkg/provider/emarsys`    | Emarsys typed events.                                                                  |
| `github.com/foomo/sesamy-go/pkg/provider/tracify`    | Tracify typed events.                                                                  |
| `github.com/foomo/sesamy-go/integration/loki`        | Push events to Grafana Loki via protobuf + snappy.                                     |

Authoritative API reference: <https://pkg.go.dev/github.com/foomo/sesamy-go>.
