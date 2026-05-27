# Events

## The Event[P] type

```go
type Event[P any] struct {
	Name   EventName `json:"name"`
	Params P         `json:"params,omitempty"`
}

func NewEvent[P any](name EventName, params P) Event[P]
```

Helpers on `Event[P]`:

| Method            | Description                                                  |
| ----------------- | ------------------------------------------------------------ |
| `AnyEvent()`      | Project to `Event[any]` for protocol-level encoding.         |
| `Decode(out)`     | Decode this event into another shape via mapstructure.       |
| `DecodeParams(out)` | Decode only `.Params`.                                     |
| `EncodeParams(in)` | Encode an arbitrary value into this event's `.Params`.      |

## Predefined GA4 event names

`pkg/sesamy/eventname.go` defines 44 constants for [GA4 automatically collected, enhanced-measurement, and recommended events](https://support.google.com/analytics/answer/9234069). Examples:

```
ad_impression, add_payment_info, add_shipping_info, add_to_cart, add_to_wishlist,
begin_checkout, campaign_details, click, close_convert_lead, close_unconvert_lead,
disqualify_lead, earn_virtual_money, exception, file_download, first_visit,
form_start, form_submit, generate_lead, join_group, level_end, level_start,
level_up, login, logout, page_view, post_score, purchase, qualify_lead, refund,
remove_from_cart, screen_view, scroll, search, select_content, select_item,
select_promotion, share, sign_up, spend_virtual_currency, tutorial_begin,
tutorial_complete, unlock_achievement, user_engagement, video_complete,
video_progress, video_start, view_cart, view_item, view_item_list, view_promotion,
view_search_results
```

Use them as the first argument to `NewEvent`:

```go
e := sesamy.NewEvent(sesamy.EventNamePurchase, params)
```

Custom event names are fine too — `EventName` is just a `string`:

```go
const MyEvent sesamy.EventName = "my_custom_event"
```

## Custom typed params

Define a struct and use the generic:

```go
type PurchaseParams struct {
	TransactionID string  `json:"transaction_id"`
	Value         float64 `json:"value"`
	Currency      string  `json:"currency"`
}

ev := sesamy.NewEvent(sesamy.EventNamePurchase, PurchaseParams{
	TransactionID: "T-123",
	Value:         42,
	Currency:      "EUR",
})

// To transport: project to any-params:
any := ev.AnyEvent()
```

## Provider-defined events

Providers ship their own typed events that wrap `sesamy.Event[P]`:

- **Emarsys** — `EmarsysCart`, `EmarsysCategory`, `EmarsysPurchase`, `EmarsysView`.
- **Tracify** — conversion-tracking events.

See [reference/providers](/reference/providers/emarsys).

## Decoding inbound events

A common pattern: react to specific event names with typed params.

```go
func handle(l *zap.Logger, r *http.Request, e *sesamy.Event[any]) error {
	switch e.Name {
	case sesamy.EventNamePurchase:
		var p PurchaseParams
		if err := e.DecodeParams(&p); err != nil {
			return err
		}
		// do something with typed p
	}
	return nil
}
```

Plug into the middleware chain via `gtaghttp.MiddlewareEventHandler` or `mpv2http.MiddlewareEventHandler` — see [server-side Collect](/guide/server-collect#cross-protocol-event-hooks).
