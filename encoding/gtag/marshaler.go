package gtag

type Marshler interface {
	MarshalMPv2() (*Payload, error)
}
