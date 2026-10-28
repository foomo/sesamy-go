package mpv2

// Consent is a Measurement Protocol consent state.
type Consent string

// Consent states.
const (
	ConsentDenied  Consent = "DENIED"
	ConsentGranted Consent = "GRANTED"
)

// ConsentText returns [ConsentGranted] if v is true and [ConsentDenied]
// otherwise.
func ConsentText(v bool) Consent {
	if v {
		return ConsentGranted
	}

	return ConsentDenied
}
