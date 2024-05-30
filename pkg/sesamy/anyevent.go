package sesamy

// AnyEvent casting is required as castings like Event[any](pageView) do not work.
type AnyEvent interface {
	AnyEvent() Event[any]
}
