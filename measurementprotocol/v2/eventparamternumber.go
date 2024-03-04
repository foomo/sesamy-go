package v2

type EventParameterNumber string

const (
	EventParameterNumberValue EventParameterNumber = "value"
	EventParameterNumberShipping EventParameterNumber = "shipping"
	EventParameterNumberTax EventParameterNumber = "tax"
)

func (s EventParameterNumber) String() string {
	return string(s)
}
