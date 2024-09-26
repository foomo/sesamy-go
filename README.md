# Sesamy Go SDK

[![Build Status](https://github.com/foomo/sesamy-go/actions/workflows/test.yml/badge.svg?branch=main&event=push)](https://github.com/foomo/sesamy-go/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/foomo/sesamy-go)](https://goreportcard.com/report/github.com/foomo/sesamy-go)
[![GoDoc](https://godoc.org/github.com/foomo/sesamy-go?status.svg)](https://godoc.org/github.com/foomo/sesamy-go)

> **Se**rver **S**ide T**a**g **M**anagement **S**ystem


## References

- [Event naming rules](https://support.google.com/analytics/answer/13316687)

## How to Contribute

Make a pull request...

## License

Distributed under MIT License, please see license file within the code for more details.

## Glossary

As some terms for Google Analytics and GTM events and ressources can be confusing we want to clarify the following definitions that are used in this README:

- Google Tag Manager: A configuration interface for different kind of containers (Web, Server etc.). The server container configuration is represented by a Tagging server
- Web container: This is a tagging container hosted by Google and can be configured via the Google Tag Manager
- Tagging server: This is a server that you setup yourself based on an docker image from Google that can receive and process "/g/collect" and other requests. You have to set environment variables to connect your server with the Google Tag Manager configuration interface
- Preview server: Similar to the tagging server, with the same docker image, but different environment variables to be needed. This server provides you with a Web UI to preview incoming and outgoing requests of the Tagging server
- "/g/collect" aka "Measurement Protocol V2" request: This is a GET request done by the gtag.js script installed on your website and going to the tagging server. Before Google Tag Manager, this request would have gone directly to Google Analytics.
- "/mp/collect" request: This is a different way of sending events from the website to the tagging server via a POST request including a JSON body defined here: https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference?client_type=gtag#payload

## How does it work

Sesamy-go is a Windmill (https://watermill.io/) powered Google Analytics web tracking event processor that decodes and unmarshals incoming measurement protocol requests on a website into a `mpv2.Event`. Those requests are sent by the gtag.js script, that is loaded via the gtm.js script installed on a website.

Sesamy comes with a subscriber for keel webserver (https://github.com/foomo/keel) and receives the incoming "/g/collect" request from your website.

In order to easier decode and transform the requests into a `mpv2.Event` struct, it comes partly with predefined types based on standard Google Analytics event types defined here: https://developers.google.com/tag-platform/gtagjs/reference/events

- add_payment_info
- add_shipping_info
- add_to_cart
- add_to_wishlist
- begin_checkout
- login
- purchase
- refund
- remove_from_cart
- search
- sign_up

You can extend the list of event types and their properties in your own application. For adding events and extending exisiting ones see [link](#adding-new-events-or-extending-existing-ones)

## Windmill integration

Windmill is a go library for building event-driven applications (https://watermill.io/). It wraps multiple different pipeline event systems like Kafka, Nats Jetstream etc. If you want to run a service as a singleton you can also use go channels.

Windmill provides a router which you can add handlers to

```go
// main.go

import (
  	"github.com/ThreeDotsLabs/watermill/message"
    "github.com/pperaltaisern/watermillzap"
)

func main () {
  ...
  router, err := message.NewRouter(message.RouterConfig{}, watermillzap.NewLogger(l))
  ...
}
```

You can add handlers via the `addHandler` function. This function takes the following parameters

- handlerName (e.g. "http")
- subscribeTopic
- subscriber
- publishTopic
- publisher
- handlerFunc

In your application you need to instantiate a http subscriber to handle the incoming "/g/collect" requests. This subscriber then decodes and unmarshals the incoming event into a `mpv2.Event` event:

```go
// main.go

import (
  watermillkeel "github.com/foomo/sesamy-go/integration/watermill/keel"
  "github.com/ThreeDotsLabs/watermill/message"
)

func main () {
  ...

  // Create pubSub
  pubSub := gochannel.NewGoChannel(
    gochannel.Config{},
    watermillzap.NewLogger(l),
  )

  // create subscriber
  subscriber := watermillkeel.NewSubscriber(l.Named(HandlerHTTP),
    watermillkeel.SubscriberWithMiddlewares(
      ...
    ),
    watermillkeel.SubscriberWithRequireEventName(),
    watermillkeel.SubscriberWithLogger(),
  )

  // add handler to publish incoming http events
  router.AddHandler("gtm", "", subscriber, "sesamy.http", pubSub, message.PassthroughHandler)

  ...
}
```

In oder to publish the received events to the tagging server (e.g. after enriching them) you need to add a publisher to the windmill router that in the end sends a "/g/collect" request with the enriched event data:

```go
// main.go

import (
	watermillgtm "github.com/foomo/sesamy-go/integration/watermill/gtm"
  "github.com/ThreeDotsLabs/watermill/message"
)

func main () {
  ...

  l := svr.Logger()

  // create gtm publisher
  gtmPublisher := watermillgtm.NewPublisher(
  l,
    gtmTaggingServiceURL(),
    watermillgtm.PublisherWithClient(keelhttp.NewHTTPClient(
      keelhttp.HTTPClientWithTelemetry(),
      keelhttp.HTTPClientWithRoundTripware(l.Named("gtm"),
        roundtripware.RequestID(),
        roundtripware.Logger(),
      ),
    )),
  )

  // add handler to publish incoming http events
  router.AddHandler("gtm", "sesamy.http", pubsub, "", gtmPublisher, message.PassthroughHandler)

  ...
}
```

## Example: Add an enrichment handler

In order to enricht event data before it is sent to the server container you can add an enrichment handler in the pipeline. With the `watermillmpv2.EventHandler` gives you access to the event objedct as well as the original message.

```go
  // main.go

package main

import (
	watermillmpv2 "github.com/foomo/sesamy-go/measurementprotocol/v2"
)

func main () {
  ...

	// add enrichment handler to enrich events within the pipeline
	router.AddHandler("enrichment", "sesamy.http", pubSub, TopicEnrichment, pubSub,
    watermillmpv2.EventHandler(func(event *mpv2.Event, msg *message.Message) error {
      switch *event.EventName {
      case mpv2.EventNameAddToCart, mpv2.EventNameRemoveFromCart:
        // enrich these events
      default:
      }
      return nil
	  }
  ))

  ...
}
```

## Adding new events or extending existing ones

If you have custom events or you are extending existing events as those mentioned in [link](#sesamy), you can do that by defining a new struct per event in your application. Every struct needs to implement the following function (e.g. for AddToCart)

```go
func (e *AddToCart) MarshalMPv2() (*mpv2.Event, error)
```

<br>

> **_IMPORTANT:_** If you adopt existing or add new events, the configuration in the Web container needs to be adopted in Google Tag Manager as well. Otherwise those events and parameters are not received in Google Analytics. These adoptations can be done manually or programmatically via the Sesamy CLI (https://github.com/foomo/sesamy-cli)

## Sesamy-gtm helm chart

https://artifacthub.io/packages/helm/foomo/sesamy-gtm

This contains the standard preview and the tagging server from Google. Additionally, if you want to introduce enrichment functionalities you also need to use the "collect" chart as a proxy between your website and the Google tag manager server container. As a consequence it needs a main.go that is than build by the "collect" helm chart. This chart exposes a "/g/collect" path so that all events triggered by your website are sent to your service instead of directly to the standard Google Tagging Container.
