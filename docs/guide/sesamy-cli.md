# Pairing with sesamy-cli

`sesamy-go` is the **runtime** half of the Sesamy stack. [`foomo/sesamy-cli`](https://github.com/foomo/sesamy-cli) is the **build-time** half: it generates the GTM (Google Tag Manager) container configuration that points `gtag.js` traffic at your `sesamy-go` endpoint instead of `www.google-analytics.com`.

You don't need `sesamy-cli` to use `sesamy-go`, but the two are designed to work together.

## What sesamy-cli does

- Reads a YAML/HCL spec describing your GA4 setup (tags, triggers, variables, custom events, consent mode).
- Generates a GTM container JSON you can import into a Google Tag Manager workspace.
- Configures `gtag.js` to send `/g/collect` requests to your first-party `sesamy-go` endpoint.
- Emits typed Go event structs you can use with `sesamy-go`'s `Event[P]` for both ingest and provider mappers.

## End-to-end flow

```
sesamy-cli  ──┐
              │ generates  ┌──────────────────┐
              ├───────────▶│ GTM container    │ ──── imported into ───▶ Google Tag Manager
              │            └──────────────────┘
              │ generates  ┌──────────────────┐
              └───────────▶│ Go event structs │ ──── go get ──▶  your project
                           └──────────────────┘                       │
                                                                      │ used by
                                                                      ▼
                                                              ┌──────────────────┐
                                              gtag.js  ──────▶│   sesamy-go      │
                                                              │   collect server │
                                                              └──────────────────┘
```

## Typical project layout

```
my-tagging-project/
├── sesamy.yaml             # sesamy-cli input
├── generated/
│   ├── container.json      # for GTM
│   └── events.go           # imports github.com/foomo/sesamy-go
├── cmd/collect/
│   └── main.go             # uses sesamy-go collect.New(...)
```

## Workflow

1. **Author** `sesamy.yaml` — declare events, parameters, consent groups.
2. **Run** `sesamy-cli generate` — produces `container.json` and Go event types under `generated/`.
3. **Import** `container.json` into your GTM workspace.
4. **Deploy** `cmd/collect` (a thin `sesamy-go` server) under your first-party domain.
5. **Publish** the GTM workspace — `gtag.js` traffic now flows to your `sesamy-go` endpoint.

## Sharing event types

Because `sesamy-cli` emits Go types built on top of `sesamy.Event[P]`, the same struct definitions are used in:

- Your `sesamy-go` middleware to assert / transform events.
- Backend code that emits events server-side via `client.MPv2`.
- Provider mappers (e.g. an Emarsys hook that fires when an `EventNamePurchase` arrives).

See `pkg/provider/emarsys/event` for an example of typed events plugging into the `Event[P]` model.

## More

- `sesamy-cli` docs: <https://github.com/foomo/sesamy-cli>
- `sesamy-go` collect setup: [server-side Collect](/guide/server-collect)
