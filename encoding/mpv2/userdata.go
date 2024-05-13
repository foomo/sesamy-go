package mpv2

// UserData https://developers.google.com/analytics/devguides/collection/ga4/uid-data
type UserData struct {
	SHA256EmailAddress []SHA256Hash      `json:"sha256_email_address,omitempty"`
	SHA256PhoneNumber  []SHA256Hash      `json:"sha256_phone_number,omitempty"`
	Address            []UserDataAddress `json:"address,omitempty"`
}
