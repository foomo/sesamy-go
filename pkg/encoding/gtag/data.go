package gtag

// Data is the intermediate map that gtag.js query parameters are decoded into
// before they are mapped onto a [Payload].
type Data map[string]any
