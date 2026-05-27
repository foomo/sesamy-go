# Emarsys

`pkg/provider/emarsys`.

Typed GA4 events for [Emarsys](https://emarsys.com/) marketing automation.

## Event names

| Constant                       | Wire name           |
| ------------------------------ | ------------------- |
| `EventNameEmarsysCart`         | `emarsys_cart`      |
| `EventNameEmarsysCategory`     | `emarsys_category`  |
| `EventNameEmarsysPurchase`     | `emarsys_purchase`  |
| `EventNameEmarsysView`         | `emarsys_view`      |

## Event constructors

Each lives in `pkg/provider/emarsys/event/`:

```go
import (
	emarsysevent "github.com/foomo/sesamy-go/pkg/provider/emarsys/event"
	"github.com/foomo/sesamy-go/pkg/provider/emarsys/event/params"
	sesamyparams "github.com/foomo/sesamy-go/pkg/event/params"
)

// Cart
ev := emarsysevent.NewEmarsysCart(params.EmarsysCart[sesamyparams.Item]{
	Items: []sesamyparams.Item{ /* ... */ },
})

// Category
evC := emarsysevent.NewEmarsysCategory(params.EmarsysCategory{Category: "shoes/boots"})

// Purchase
evP := emarsysevent.NewEmarsysPurchase(params.EmarsysPurchase[sesamyparams.Item]{
	TransactionID: "T-123",
	Items: []sesamyparams.Item{ /* ... */ },
})

// View
evV := emarsysevent.NewEmarsysView(params.EmarsysView[sesamyparams.Item]{
	Items: []sesamyparams.Item{ /* ... */ },
})
```

All four return `sesamy.Event[P]` and can be projected to `Event[any]` via `.AnyEvent()` for transport.

## Params

`pkg/provider/emarsys/event/params` defines the typed param structs. They are generic over an `Item` type so that you can reuse `sesamyparams.Item` (the canonical GA4 item) or supply your own enriched item shape.

## Reacting to inbound Emarsys events

Inside a cross-protocol `EventHandler`:

```go
func handle(l *zap.Logger, r *http.Request, e *sesamy.Event[any]) error {
	switch e.Name {
	case emarsysevent.EventNameEmarsysPurchase:
		var p params.EmarsysPurchase[sesamyparams.Item]
		if err := e.DecodeParams(&p); err != nil {
			return err
		}
		// forward p to Emarsys' Web Behavior API ...
	}
	return nil
}
```
