package cookiebot

const CookieName = "CookieConsent"

type Cookie struct {
	Stamp       string `json:"stamp"`
	Necessary   bool   `json:"necessary"`
	Preferences bool   `json:"preferences"`
	Statistics  bool   `json:"statistics"`
	Marketing   bool   `json:"marketing"`
	Method      string `json:"method"`
	Version     string `json:"ver"`
	UTC         int    `json:"utc"`
	Region      string `json:"region"`
}
