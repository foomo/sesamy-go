package gtag

type ClientHints struct {
	// Is a random hash generated on the page load.
	// Examole: 456193680
	// RandomPageLoadHash *string `json:"random_page_load_hash,omitempty" gtag:"_p,omitempty"`
	// Browser screen resolution in format width x height
	// Example: 2560x1440
	ScreenResolution *string `json:"screen_resolution,omitempty" gtag:"sr,omitempty"`
	// Browser active locale.
	// Example: es-es
	UserLanguage *string `json:"user_language,omitempty" gtag:"ul,omitempty"`
	// Example: x86
	UserAgentArchitecture *string `json:"user_agent_architecture,omitempty" gtag:"uaa,omitempty"`
	// The "bitness" of the user-agent's underlying CPU architecture. This is the size in bits of an integer or memory address—typically 64 or 32 bits.
	// Example: 32 | 64
	UserAgentBitness *string `json:"user_agent_bitness,omitempty" gtag:"uab,omitempty"`
	// The brand and full version information for each brand associated with the browser, in a comma-separated list
	// Example: Google Chrome;105.0.5195.127|Not)A;Brand;8.0.0.0|Chromium;105.0.5195.127
	UserAgentFullVersionList *string `json:"user_agent_full_version_list,omitempty" gtag:"uafvl,omitempty"`
	// Indicates whether the browser is on a mobile device
	// Example: 1
	UserAgentMobile *string `json:"user_agent_mobile,omitempty" gtag:"uamb,omitempty"`
	// The device model on which the browser is running. Will likely be empty for desktop browsers
	// Example: Nexus 6
	UserAgentModel *string `json:"user_agent_model,omitempty" gtag:"uam,omitempty"`
	// The platform or operating system on which the user agent is running
	// Example: Chromium OS | macOS | Android | iOS
	UserAgentPlatform *string `json:"user_agent_platform,omitempty" gtag:"uap,omitempty"`
	// The version of the operating system on which the user agent is running
	// Example: 14.0.0
	UserAgentPlatformVersion *string `json:"user_agent_platform_version,omitempty" gtag:"uapv,omitempty"`
	// Whatever Windows On Windows 64 Bit is supported. Used by "WoW64-ness" sites. ( running 32bits app on 64bits windows)
	// Example: 1
	UserAgentWOW64 *string `json:"user_agent_wow_64,omitempty" gtag:"uaw,omitempty"`
	// Added to report the current country for the user under some circumstanced. To be documented.
	// Example: ES
	UserCountry *string `json:"user_country,omitempty" gtag:"_uc,omitempty"`
	// Example: DE-BY
	UserRegion *string `json:"user_region,omitempty" gtag:"ur,omitempty"`
}
