package mpv2

type Consent string

const (
	ConsentDenied  Consent = "DENIED"
	ConsentGranted Consent = "GRANTED"
)

func ConsentText(v bool) Consent {
	if v {
		return ConsentGranted
	}
	return ConsentDenied
}
