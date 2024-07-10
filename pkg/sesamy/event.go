package sesamy

type Event[P any] struct {
	// Reserved names: https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference?client_type=gtag#reserved_names
	Name EventName `json:"name"`
	// Reserved parameter names: https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference?client_type=gtag#reserved_parameter_names
	Params P `json:"params,omitempty"`
}

// ------------------------------------------------------------------------------------------------
// ~ Constructor
// ------------------------------------------------------------------------------------------------

func NewEvent[P any](name EventName, params P) Event[P] {
	return Event[P]{
		Name:   name,
		Params: params,
	}
}

// ------------------------------------------------------------------------------------------------
// ~ Public methods
// ------------------------------------------------------------------------------------------------

func (e Event[P]) AnyEvent() Event[any] {
	return Event[any]{
		Name:   e.Name,
		Params: e.Params,
	}
}

func (e Event[P]) Decode(output any) error {
	return Decode(e, output)
}

func (e Event[P]) DecodeParams(output any) error {
	return Decode(e.Params, output)
}
