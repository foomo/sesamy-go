// Package client sends GA4 events to a server-side tagging endpoint using the
// gtag.js protocol ([GTag]) or Measurement Protocol v2 ([MPv2]).
//
// Both clients run payloads through a middleware chain that fills in IDs from
// the GA cookies of the incoming request before sending. Only selected cookies
// of the incoming request, by default gtm_auth, gtm_debug and gtm_preview, are
// forwarded upstream.
package client
