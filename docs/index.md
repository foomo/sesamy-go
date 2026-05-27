---
layout: home

hero:
  name: "sesamy-go"
  text: "Server-side tag management System for Go"
  tagline: GA4 tracking via GTag and Measurement Protocol v2, with first-class server-side support.
  image:
    src: /logo.png
    alt: sesamy-go
  actions:
    - theme: brand
      text: Get Started
      link: /guide/getting-started
    - theme: alt
      text: What is sesamy-go?
      link: /guide/what-is-sesamy
    - theme: alt
      text: GitHub
      link: https://github.com/foomo/sesamy-go

features:
  - title: GTag + MPv2
    details: Encode, decode, and convert between GA4's gtag.js form-encoded protocol and the JSON-based Measurement Protocol v2.
  - title: Server-side Collect
    details: Drop-in HTTP handlers that receive client tracking requests and forward them to a server-side tagging endpoint.
  - title: Composable middleware
    details: Stack-friendly middleware for client ID, session ID, user ID, page location, IP override, engagement time, debug mode, logging, and tracing.
  - title: Typed events
    details: Generic Event[P] with 44 predefined GA4 event names and typed parameter structs.
  - title: Providers built-in
    details: Cookiebot consent parsing, Emarsys marketing events, Tracify conversion tracking.
  - title: Loki integration
    details: Ship events to Grafana Loki via protobuf with snappy compression and batched delivery.
  - title: Pairs with sesamy-cli
    details: Generates GTM container configs that route gtag.js traffic through your sesamy-go server endpoint.
  - title: OpenTelemetry-aware
    details: Trace IDs are attached to log lines when an active span is present.
---
