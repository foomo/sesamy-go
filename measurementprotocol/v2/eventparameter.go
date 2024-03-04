package v2

type EventParameter string

const (
	EventParameterMethod EventParameter = "method"
	EventParameterCoupon EventParameter = "coupon"
	EventParameterPaymentType EventParameter = "payment_type"
	EventParameterShippingTier EventParameter = "shipping_tier"
	EventParameterTransactionID EventParameter = "transaction_id"
	EventParameterSearchTerm EventParameter = "search_term"
)

func (s EventParameter) String() string {
	return string(s)
}

