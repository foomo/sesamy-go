// Package loki ships GA4 events to Grafana Loki. [Loki] buffers events, batches
// them and pushes them as snappy-compressed protobuf with exponential backoff
// on failure. [GTagMiddleware] and [MPv2Middleware] feed it from the HTTP
// handlers, and [Service] runs it as a service.
package loki
