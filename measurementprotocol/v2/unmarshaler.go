package v2

type Unmarshler interface {
	UnmarshalMPv2(*Event) error
}
