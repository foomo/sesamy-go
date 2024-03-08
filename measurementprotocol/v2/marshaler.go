package v2

type Marshler interface {
	MarshalMPv2() (*Event, error)
}
