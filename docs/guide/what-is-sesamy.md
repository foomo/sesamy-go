# What is sesamy-go?

`sesamy-go` is a Go SDK for **server-side tag management (SSTM)** targeted at Google Analytics 4 (GA4). It gives you the building blocks to receive, transform, enrich, and forward GA4 tracking traffic on your own infrastructure instead of sending it directly from the browser to Google.

## Why server-side?

Browser-side tracking has well-known problems:

- Ad-blockers and tracking-protection break event collection.
- First-party cookies set by third-party scripts are short-lived (Safari ITP, Chrome 3rd-party cookie deprecation).
- Sensitive parameters (IP, user-agent, user IDs, marketing identifiers) leak to vendor endpoints uncontrolled.
- You can't normalize, validate, redact, or enrich payloads before they hit GA.

Server-side tagging solves these by collecting events on a **first-party endpoint** under your domain. `sesamy-go` is the SDK to build that endpoint in Go.

## What it does

- **Receives** GA4 traffic in either of the two protocols GA4 uses:
  - **GTag** — the form-encoded query-string protocol used by `gtag.js`. See [reference/encoding](/reference/encoding).
  - **Measurement Protocol v2 (MPv2)** — JSON-based, used by SDK-style senders.
- **Parses** payloads into Go structs you can read, modify, and assert against.
- **Converts** between GTag ↔ MPv2 so middleware can be written once against MPv2 events even when traffic arrives as GTag.
- **Enriches** events with client ID, session ID, user ID, IP override, page location, engagement time, user-agent — via composable middleware.
- **Forwards** enriched events to a downstream tagging server (Server-side GTM, BigQuery sink, vendor endpoints) or directly to GA4.
- **Ships** server-side events to Grafana Loki for observability ([integration/loki](/guide/loki)).
- **Integrates** with consent managers (Cookiebot) and marketing platforms (Emarsys, Tracify).

## What it is not

- Not a hosted SaaS. You run it.
- Not a GTM replacement. It pairs with [`foomo/sesamy-cli`](/guide/sesamy-cli), which generates the GTM container that points `gtag.js` at your `sesamy-go` endpoint.
- Not a generic web framework. It exposes `http.Handler` functions; bring your own router.

## How it fits together

```
┌───────────┐     gtag.js / mpv2     ┌─────────────────┐    forwarded     ┌──────────────┐
│  browser  │ ──────────────────────▶│   sesamy-go     │ ────────────────▶│ GA4 / sGTM   │
└───────────┘  (your first-party     │   collect       │  enriched +      │ / BigQuery   │
               domain)               │   middleware    │  normalized      └──────────────┘
                                     └─────────────────┘
                                              │
                                              ▼
                                     ┌─────────────────┐
                                     │ Loki / logs /   │
                                     │ providers       │
                                     └─────────────────┘
```

Next: [Getting started](/guide/getting-started).
