package cookiebot

const CookieName = "CookieConsent"

// {stamp:'VLZnHUKBPLqZCJyClLLmnGglmUPeZsGxrmiAEZ48i7UH39ptKHY4MA==',necessary:true,preferences:true,statistics:true,marketing:true,method:'explicit',ver:1,utc:1724770548958,region:'de'}
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
