package cookiebot

const CookieName = "CookieConsent"

type Cookie struct {
	Stamp       string `json:"stamp" yaml:"stamp"`
	Necessary   bool   `json:"necessary" yaml:"necessary"`
	Preferences bool   `json:"preferences" yaml:"preferences"`
	Statistics  bool   `json:"statistics" yaml:"statistics"`
	Marketing   bool   `json:"marketing" yaml:"marketing"`
	Method      string `json:"method" yaml:"method"`
	Version     string `json:"ver" yaml:"ver"`
	UTC         int    `json:"utc" yaml:"utc"`
	Region      string `json:"region" yaml:"region"`
}
