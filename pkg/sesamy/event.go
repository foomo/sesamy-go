package sesamy

// Event is a GA4 event with name Name and parameters of type P.
// It encodes to the Measurement Protocol event JSON shape.
type Event[P any] struct {
	// Reserved names: https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference?client_type=gtag#reserved_names
	Name EventName `json:"name"`
	// Reserved parameter names: https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference?client_type=gtag#reserved_parameter_names
	Params P `json:"params,omitempty"`
}

// ------------------------------------------------------------------------------------------------
// ~ Constructor
// ------------------------------------------------------------------------------------------------

// NewEvent returns an Event with the given name and params.
func NewEvent[P any](name EventName, params P) Event[P] {
	return Event[P]{
		Name:   name,
		Params: params,
	}
}

// ------------------------------------------------------------------------------------------------
// ~ Public methods
// ------------------------------------------------------------------------------------------------

// AnyEvent returns a copy of e with its params typed as any.
func (e Event[P]) AnyEvent() Event[any] {
	return Event[any]{
		Name:   e.Name,
		Params: e.Params,
	}
}

// Decode decodes e into output, which must be a pointer, using [Decode].
func (e Event[P]) Decode(output any) error {
	return Decode(e, output)
}

// DecodeParams decodes the params of e into output, which must be a pointer,
// using [Decode].
func (e Event[P]) DecodeParams(output any) error {
	return Decode(e.Params, output)
}

// EncodeParams decodes input into the params of e using [Decode].
// Because e is passed by value, the change is only visible to the caller
// when P is a reference type such as a map.
func (e Event[P]) EncodeParams(input any) error {
	return Decode(input, &e.Params)
}
