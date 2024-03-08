package v2

// EventParameterNumber as number
// See https://support.google.com/analytics/table/13594742?sjid=7861230991468479976-EU
type EventParameterNumber string

const (
	EventParameterNumberValue EventParameterNumber = "value"
	// EventParameterNumberFirebaseError The error code reported by the Firebase SDK.
	EventParameterNumberFirebaseError EventParameterNumber = "firebase_error"
	// EventParameterNumberFreeTrial Signifies that an in-app purchase is a free trial
	EventParameterNumberFreeTrial EventParameterNumber = "free_trial"
	// EventParameterNumberMessageDeviceTime The Firebase Cloud Messaging or Firebase In-App Messaging delivery epoch timestamp in UTC.	(None)
	EventParameterNumberMessageDeviceTime EventParameterNumber = "message_device_time"
	// EventParameterNumberMessageTime The Firebase Cloud Messaging message notification epoch timestamp in UTC.	(None)
	EventParameterNumberMessageTime EventParameterNumber = "message_time"
	// EventParameterNumberShipping The shipping cost associated with a transaction
	EventParameterNumberShipping EventParameterNumber = "shipping"
	// EventParameterNumberTax The tax cost associated with a transaction
	EventParameterNumberTax EventParameterNumber = "tax"
)

func (s EventParameterNumber) String() string {
	return string(s)
}
