package mpv2

// UserDataAddress https://developers.google.com/analytics/devguides/collection/ga4/uid-data
type UserDataAddress struct {
	SHA256FirstName SHA256Hash `json:"sha256_first_name,omitempty"`
	SHA256LastName  SHA256Hash `json:"sha256_last_name,omitempty"`
	SHA256Street    SHA256Hash `json:"sha256_street,omitempty"`
	City            string     `json:"city,omitempty"`
	Region          string     `json:"region,omitempty"`
	PostalCode      string     `json:"postal_code,omitempty"`
	Country         string     `json:"country,omitempty"`
}
