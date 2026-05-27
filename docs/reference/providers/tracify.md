# Tracify

`pkg/provider/tracify`.

Typed events for [Tracify](https://www.tracify.ai/) attribution / conversion tracking. Event constructors live under `pkg/provider/tracify/event/` and follow the same pattern as the Emarsys provider:

- Define `const EventNameTracify… sesamy.EventName = "..."`.
- Provide `NewTracify…(params...) sesamy.Event[P]` constructors.
- Use typed param structs under `pkg/provider/tracify/event/params/`.

Refer to the package source for the current set of constructors and params. The same usage shape applies as in [Emarsys](/reference/providers/emarsys):

```go
import (
	tracifyevent "github.com/foomo/sesamy-go/pkg/provider/tracify/event"
)

ev := tracifyevent.NewTracifySomething(/* typed params */)
any := ev.AnyEvent()
```

And reactively:

```go
func handle(l *zap.Logger, r *http.Request, e *sesamy.Event[any]) error {
	if e.Name == tracifyevent.EventNameTracifySomething {
		var p params.TracifySomething
		if err := e.DecodeParams(&p); err != nil {
			return err
		}
		// forward to Tracify ...
	}
	return nil
}
```
