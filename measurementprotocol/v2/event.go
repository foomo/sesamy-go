package v2

// See https://www.thyngster.com/ga4-measurement-protocol-cheatsheet/
type Event struct {
	// --- Request parameters ---

	// Defines que current protocol version being used.
	// Example: 2
	ProtocolVersion *string `json:"v,omitempty" mapstructure:"v,omitempty"`
	// Current Stream ID / Measurement ID
	// Example: G-THYNGSTER
	TrackingID *string `json:"tid,omitempty" mapstructure:"tid,omitempty"`
	// If the current hit is coming was generated from GTM, it will contain a hash of current GTM/GTAG config
	// Example: 2oear0
	GTMHashInfo *string `json:"gtm,omitempty" mapstructure:"gtm,omitempty"`
	// Is a random hash generated on the page load.
	// Examole: 456193680
	RandomPageLoadHash *string `json:"_p,omitempty" mapstructure:"_p,omitempty"`
	// Browser screen resolution in format width x height
	// Example: 2560x1440
	ScreenResolution *string `json:"sr,omitempty" mapstructure:"sr,omitempty"`
	// Browser active locale.
	// Example: es-es
	UserLanguage *string `json:"ul,omitempty" mapstructure:"ul,omitempty"`
	// Current Document Hostname
	// Exampple: www.analytics-debugger.com
	DocumentHostname *string `json:"dh,omitempty" mapstructure:"dh,omitempty"`
	// Google Analytics Client Id
	// Example: 281344611.1635634925
	ClientID *string `json:"cid,omitempty" mapstructure:"cid,omitempty"`
	// Current hits counter for the current page load
	// Example: 1
	HitCounter *string `json:"_s,omitempty" mapstructure:"_s,omitempty"`
	// This is supposed to be to enrich the GA4 hits to send data to SGTM, at this point is always set as an empty value...
	Richsstsse *string `json:"richsstsse,omitempty" mapstructure:"richsstsse,omitempty"`

	// --- Client Hints ---

	// Example: x86
	UserAgentArchitecture *string `json:"uaa,omitempty" mapstructure:"uaa,omitempty"`
	// The "bitness" of the user-agent's underlying CPU architecture. This is the size in bits of an integer or memory address—typically 64 or 32 bits.
	// Example: 32 | 64
	UserAgentBitness *string `json:"uab,omitempty" mapstructure:"uab,omitempty"`
	// The brand and full version information for each brand associated with the browser, in a comma-separated list
	// Example: Google Chrome;105.0.5195.127|Not)A;Brand;8.0.0.0|Chromium;105.0.5195.127
	UserAgentFullVersionList *string `json:"uafvl,omitempty" mapstructure:"uafvl,omitempty"`
	// Indicates whether the browser is on a mobile device
	// Example: 1
	UserAgentMobile *string `json:"uamb,omitempty" mapstructure:"uamb,omitempty"`
	// The device model on which the browser is running. Will likely be empty for desktop browsers
	// Example: Nexus 6
	UserAgentModel *string `json:"uam,omitempty" mapstructure:"uam,omitempty"`
	// The platform or operating system on which the user agent is running
	// Example: Chromium OS | macOS | Android | iOS
	UserAgentPlatform *string `json:"uap,omitempty" mapstructure:"uap,omitempty"`
	// The version of the operating system on which the user agent is running
	// Example: 14.0.0
	UserAgentPlatformVersion *string `json:"uapv,omitempty" mapstructure:"uapv,omitempty"`
	// Whatever Windows On Windows 64 Bit is supported. Used by "WoW64-ness" sites. ( running 32bits app on 64bits windows)
	// Example: 1
	UserAgentWOW64 *string `json:"uaw,omitempty" mapstructure:"uaw,omitempty"`

	// --- Shared ---

	// Actual page's Pathname. It does not include the hostname, quertyString or Fragment
	// Example: /hire-me
	DocumentLocation *string `json:"dl,omitempty" mapstructure:"dl,omitempty"`
	// Actual page's Title
	// Example: Hire Me
	DocumentTitle *string `json:"dt,omitempty" mapstructure:"dt,omitempty"`
	// Actual page's Referrer
	// Example:
	DocumentReferrer *string `json:"dr,omitempty" mapstructure:"dr,omitempty"`
	// Unknown. Value ccd.{{HASH}}. The hash in based on various internal parameters. Some kind of usage hash.
	// Example: ccd.AAB
	Z *string `json:"_z,omitempty" mapstructure:"_z,omitempty"`
	// This is added when an event is generated from rules (from the admin). Actually is hash of the "GA4_EVENT" string
	// Example: Q
	EventUsage *string `json:"_eu,omitempty" mapstructure:"_eu,omitempty"`
	// Unknown
	// Example:
	EventDebugID *string `json:"edid,omitempty" mapstructure:"edid,omitempty"`
	// If an event contains this parameters it won't be processed and it will show on on the debug View in GA4
	// Example: 1
	IsDebug *string `json:"_dbg,omitempty" mapstructure:"_dbg,omitempty"`
	// If the current request has a referrer, it will be ignored at processing level
	// Example: 1
	IgnoreReferrer *string `json:"ir,omitempty" mapstructure:"ir,omitempty"`
	// Traffic Type
	// Example: 1
	TrafficType *string `json:"tt,omitempty" mapstructure:"tt,omitempty"`
	// Current Google Consent Status. Format 'G1'+'AdsStorageBoolStatus'`+'AnalyticsStorageBoolStatus'
	// Example:  G101
	GoogleConsentStatus *string `json:"gcs,omitempty" mapstructure:"gcs,omitempty"`
	// Will be added with the value "1" if the Google Consent has just been updated (wait_for_update setting on GTAG)
	// Example: 1
	GoogleConsentUpdate *string `json:"gcu,omitempty" mapstructure:"gcu,omitempty"`
	// Documented values, 1 or 2, no more info on the meaning
	// Example: 2
	GoogleConsentUpdateType *string `json:"gcut,omitempty" mapstructure:"gcut,omitempty"`
	// Will be added with the value "1" if the Google Consent had a default value before getting an update
	// Example: G111
	GoogleConsentDefault *string `json:"gcd,omitempty" mapstructure:"gcd,omitempty"`
	// Will be set to 1 is the current page has a linker and this last one is valid
	// Example: 1
	IsGoogleLinkerValid *string `json:"_glv,omitempty" mapstructure:"_glv,omitempty"`

	// --- Campaign Attributes ---

	// Campaign Medium ( utm_medium ), this will override the current values read from the url
	// Example: cpc
	CampaignMedium *string `json:"cm,omitempty" mapstructure:"cm,omitempty"`
	// Campaign Source ( utm_source ), this will override the current values read from the url
	// Example: google
	CampaignSource *string `json:"cs,omitempty" mapstructure:"cs,omitempty"`
	// Campaign Name ( utm_campaign ), this will override the current values read from the url
	// Example: cpc
	CampaignName *string `json:"cn,omitempty" mapstructure:"cn,omitempty"`
	// Campaign Content ( utm_content ), this will override the current values read from the url
	// Example: big banner
	CampaignContent *string `json:"cc,omitempty" mapstructure:"cc,omitempty"`
	// Campaign Term ( utm_term ), this will override the current values read from the url
	// Example: summer
	CampaignTerm *string `json:"ck,omitempty" mapstructure:"ck,omitempty"`
	// Campaign Creative Format ( utm_creative_format ), this will override the current values read from the url
	// Example: native
	CampaignCreativeFormat *string `json:"ccf,omitempty" mapstructure:"ccf,omitempty"`
	// Campaign Marketing Tactic ( utm_marketing_tactic ), this will override the current values read from the url
	// Example: prospecting
	CampaignMarketingTactic *string `json:"cmt,omitempty" mapstructure:"cmt,omitempty"`
	// Random Number used to Dedupe gclid
	// Example: 342342343
	GclidDeduper *string `json:"_rnd,omitempty" mapstructure:"_rnd,omitempty"`

	// --- Event Parameters ---

	// Current Event Name.
	// Example: page_view
	EventName *EventName `json:"en,omitempty" mapstructure:"en,omitempty"`
	// It's the total engagement time in milliseconds since the last event. The engagement time is measured only when the current page is visible and active ( ie: the browser window/tab must be active and visible ), for this GA4 uses the window.events: focus, blur, pageshow, pagehide and the document:visibilitychange, these will determine when the timer starts and pauses
	// Example: 1234
	EngagementTime *string `json:"_et,omitempty" mapstructure:"_et,omitempty"`
	// Defines a parameter for the current Event
	// Example: ep.page_type: checkout
	EventParameter map[string]string `json:"ep,omitempty" mapstructure:"ep,omitempty"`
	// Defines a parameter for the current Event
	// Example: epn.plays_count: 42
	EventParameterNumber map[string]string `json:"epn,omitempty" mapstructure:"epn,omitempty"`
	// If the current event is set as a conversion on the admin interacted the evfent will have this value present
	// Example: 1
	IsConversion *string `json:"_c,omitempty" mapstructure:"_c,omitempty"`
	// External Event
	ExternalEvent *string `json:"_ee,omitempty" mapstructure:"_ee,omitempty"`

	// --- Session / User Related ---

	// Current User ID
	// Example: 1635691016
	UserID *string `json:"uid,omitempty" mapstructure:"uid,omitempty"`
	// Current Firebase ID
	// Example: HASHSAH
	FirebaseID *string `json:"_fid,omitempty" mapstructure:"_fid,omitempty"`
	// GA4 Session Id. This comes from the GA4 Cookie. It may be different for each Stream ID Configured on the site
	// Example: 1635691016
	SessionID *string `json:"sid,omitempty" mapstructure:"sid,omitempty"`
	// Count of sessions recorded by GA4. This value increases by one each time a new session is detected ( when the session expires )
	// Example: 10
	SessionCount *string `json:"sct,omitempty" mapstructure:"sct,omitempty"`
	// If the current user is engaged in any way, this value will be 1
	// Example:
	SessionEngagment *string `json:"seg,omitempty" mapstructure:"seg,omitempty"`
	// Defines an user Propery for the current Measurement ID
	// Example: up.is_premium_user: yes
	UserProperty map[string]string `json:"up,omitempty" mapstructure:"up,omitempty"`
	// Defines an user Propery for the current Measurement ID
	// Example:
	UserPropertyNumber map[string]string `json:"upn,omitempty" mapstructure:"upn,omitempty"`
	// If the "_ga_THYNGSTER" cookie is not set, the first event will have this value present. This will internally create a new "first_visit" event on GA4. If this event is also a conversion the value will be "2" if not, will be "1"
	// Example: 1|2
	FirstVisit *string `json:"_fv,omitempty" mapstructure:"_fv,omitempty"`
	// If the "_ga_THYNGSTER" cookie last session time value is older than 1800 seconds, the current event will have this value present. This will internally create a new "session_start" event on GA4. If this event is also a conversion the value will be "2" if not, will be "1"
	// Example: 1|2
	SessionStart *string `json:"_ss,omitempty" mapstructure:"_ss,omitempty"`
	// This seems to be related to the ServerSide hits, it's 0 if the FPLC Cookie is not present and to the current value if it's coming from a Cross Domain linker
	// Example: bVhVicbfiSXaGNxeawKaPlDQc9QXPD6bKcsn36Elden6wZNb7Q5X1iXlkTVP5iP3H3y76cgM3UIgHCaRsYfPoyLGlbiIYMPRjvnUU7KWbdWLagodzxjrlPnvaRZJkw
	FirstPartyLinkerCookie *string `json:"_fplc,omitempty" mapstructure:"_fplc,omitempty"`
	// If the current user has a GA4 session cookie, but not a GA (_ga) client id cookie, this parameter will be added to the hit
	// Example: 1
	NewSessionID *string `json:"_nsi,omitempty" mapstructure:"_nsi,omitempty"`
	// You may find this parameter if using some vendor plugin o platform ( ie: using shopify integration or a prestashop plugin )
	// Example: jdhsd87
	GoogleDeveloperID *string `json:"gdid,omitempty" mapstructure:"gdid,omitempty"`
	// Added to report the current country for the user under some circumstanced. To be documented.
	// Example: ES
	UserCountry *string `json:"_uc,omitempty" mapstructure:"_uc,omitempty"`
	// Example: DE-BY
	UserRegion *string `json:"ur,omitempty" mapstructure:"ur,omitempty"`

	// --- eCommerce ---

	// Currency Code. ISO 4217
	// Example: JPY
	Currency *string `json:"cu,omitempty" mapstructure:"cu,omitempty"`
	// Example:
	Items []*Item `json:"pr,omitempty" mapstructure:"pr,omitempty"`
	// Promotion Impression/Click Tracking. Promotion Id
	// Example: summer-offer
	PromotionID *string `json:"pi,omitempty" mapstructure:"pi,omitempty"`
	// Promotion Impression/Click Tracking. Promotion Name
	// Example: summer-offer
	PromotionName *string `json:"pn,omitempty" mapstructure:"pn,omitempty"`
	// Promotion Impression/Click Tracking. Creative Name
	// Example: red-car
	// CreativeName *string `json:"cn,omitempty" mapstructure:"cn,omitempty"`
	// Promotion Impression/Click Tracking. Promotion Slot / Position
	// Example: slide-3
	// CreativeSlot *string `json:"cs,omitempty" mapstructure:"cs,omitempty"`
	// Google Place ID: Refer to: https://developers.google.com/maps/documentation/places/web-service/place-id . Seems to be inherited from Firebase, not sure about the current use on GA4
	// Example: ChIJiyj437sx3YAR9kUWC8QkLzQ
	LocationID *string `json:"lo,omitempty" mapstructure:"lo,omitempty"`

	// --- Uncategorized / Missing Info ---
	// Example: 1
	GTMUp *string `json:"gtm_up,omitempty" mapstructure:"gtm_up,omitempty"`
	// Documented values, 1,2,3: Not sure when it's added.
	EuropeanConsentModeEnabledID *string `json:"_ecid,omitempty" mapstructure:"_ecid,omitempty"`
	// Example:
	UEI *string `json:"_uei,omitempty" mapstructure:"_uei,omitempty"`
	// It's set when a Google Join is created/imported. Google Signals
	// Example: 1
	CreateGoogleJoin *string `json:"_gaz,omitempty" mapstructure:"_gaz,omitempty"`
	// Example: Redact Device Info. Need Investigation about functionality
	RedactDeviceInfo *string `json:"_rdi,omitempty" mapstructure:"_rdi,omitempty"`
	// Geo Granularity. Need Investigation about functionality
	GeoGranularity *string `json:"_geo,omitempty" mapstructure:"_geo,omitempty"`
	// Sent on sites that implement the US Privacy User Signal Mechanism, sent if window.__uspapi is present and returning a value.
	// Example: 1YNY
	USPrivacySignal *string `json:"us_privacy,omitempty" mapstructure:"us_privacy,omitempty"`
	// Sent on sites that implements IAB GDPR-Transparency-and-Consent-Framework( TCFv2 ) Mechanism. sent if window.__tcfapi is present and returning a valid value.
	// Example: 1
	GDPR *string `json:"gdpr,omitempty" mapstructure:"gdpr,omitempty"`
	// Sent on sites that implements IAB GDPR-Transparency-and-Consent-Framework( TCFv2 ) Mechanism. sent if window.__tcfapi is present and returning a valid value.
	// Example: CPfPdAAPfPdAAAHABBENCgCsAP_AAAAAAAAAI_tf_X__b3_j-_5___t0eY1f9_7__-0zjhfdl-8N3f_X_L8X_2M7vF36tq4KuR4Eu3LBIQdlHOHcTUmw6okVrzPsbk2cr7NKJ7PEmnMbeydYGH9_n1_z-ZKY7_____77__-____3_____-_f___5_3____f_V__97fn9_____9_P___9v__9__________3___gAAAJJQAYAAgj-GgAwABBH8VABgACCP5SADAAEEfx0AGAAII_kIAMAAQR_CQAYAAgj-IgAwABBH8ZABgACCP4A.f_gAAAAAAAAA
	GDPRConsent *string `json:"gdpr_consent,omitempty" mapstructure:"gdpr_consent,omitempty"`
	// Example: sypham
	NonPersonalizedAds *string `json:"npa,omitempty" mapstructure:"npa,omitempty"`
	// Example: 1
	ARE *string `json:"are,omitempty" mapstructure:"are,omitempty"`
	// Example: 1
	DigitalMarketAct *string `json:"dma,omitempty" mapstructure:"dma,omitempty"`
	// Example: sypham
	DigitalMarketActParameters *string `json:"dma_cps,omitempty" mapstructure:"dma_cps,omitempty"`
	// Example: noapi | denied
	PrivacySandboxCookieDeprecationLabel *string `json:"pscdl,omitempty" mapstructure:"pscdl,omitempty"`
	// A timestamp measuring the difference between the moment this parameter gets populated and the moment the navigation started on that particular page.
	TFD *string `json:"tfd,omitempty" mapstructure:"tfd,omitempty"`
	SST *SST    `json:"sst,omitempty" mapstructure:"sst,omitempty"`
	PAE *string `json:"pae,omitempty" mapstructure:"pae,omitempty"`

	// --- Unresolved ---

	Unknown map[string]any `json:"-" mapstructure:",remain"`
}
