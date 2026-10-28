// Package collect provides HTTP handlers that receive gtag.js and Measurement
// Protocol v2 requests, run them through configurable middlewares and forward
// them to a server-side tagging endpoint.
//
// # Usage
//
//	c, err := collect.New(l,
//		collect.WithTagging("https://sst.example.com"),
//		collect.WithGTagHTTPMiddlewares(gtaghttp.MiddlewareLogger),
//	)
//	if err != nil {
//		return err
//	}
//	mux.HandleFunc("/g/collect", c.GTagHTTPHandler)
//	mux.HandleFunc("/mp/collect", c.MPv2HTTPHandler)
package collect
