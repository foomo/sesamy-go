package gtag

import (
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// Payload https://www.thyngster.com/ga4-measurement-protocol-cheatsheet/
type Payload struct {
	Consent     `json:"consent" gtag:",inline,squash"`
	Campaign    `json:"campaign" gtag:",inline,squash"`
	ECommerce   `json:"ecommerce" gtag:",inline,squash"`
	ClientHints `json:"client_hints" gtag:",inline,squash"`

	// --- Request parameters ---

	// Defines que current protocol version being used.
	// Example: 2
	ProtocolVersion *string `json:"protocol_version,omitempty" gtag:"v,omitempty"`
	// Current Stream ID / Measurement ID
	// Example: G-THYNGSTER
	TrackingID *string `json:"tracking_id,omitempty" gtag:"tid,omitempty"`
	// If the current hit is coming was generated from GTM, it will contain a hash of current GTM/GTAG config
	// Example: 2oear0
	GTMHashInfo *string `json:"gtmhash_info,omitempty" gtag:"gtm,omitempty"`
	// Current Document Hostname
	// Exampple: www.analytics-debugger.com
	// DocumentHostname *string `json:"document_hostname,omitempty" gtag:"dh,omitempty"`
	// Google Analytics Client Id
	// Example: 281344611.1635634925
	ClientID *string `json:"client_id,omitempty" gtag:"cid,omitempty"`
	// Current hits counter for the current page load
	// Example: 1
	// HitCounter *string `json:"hit_counter,omitempty" gtag:"_s,omitempty"`
	// This is supposed to be to enrich the GA4 hits to send data to SGTM, at this point is always set as an empty value...
	Richsstsse *string `json:"richsstsse,omitempty" gtag:"richsstsse,omitempty"`

	// --- Shared ---

	// Actual page's Pathname. It does not include the hostname, quertyString or Fragment
	// Example: /hire-me
	DocumentLocation *string `json:"document_location,omitempty" gtag:"dl,omitempty"`
	// Actual page's Title
	// Example: Hire Me
	DocumentTitle *string `json:"document_title,omitempty" gtag:"dt,omitempty"`
	// Actual page's Referrer
	// Example:
	DocumentReferrer *string `json:"document_referrer,omitempty" gtag:"dr,omitempty"`
	// Unknown. Value ccd.{{HASH}}. The hash in based on various internal parameters. Some kind of usage hash.
	// Example: ccd.AAB
	// Z *string `json:"z,omitempty" gtag:"_z,omitempty"`
	// This is added when an event is generated from rules (from the admin). Actually is hash of the "GA4_EVENT" string
	// Example: Q
	// EventUsage *string `json:"event_usage,omitempty" gtag:"_eu,omitempty"`
	// Unknown
	// Example:
	// EventDebugID *string `json:"event_debug_id,omitempty" gtag:"edid,omitempty"`
	// If an event contains this parameters it won't be processed and it will show on on the debug View in GA4
	// Example: 1
	IsDebug *string `json:"is_debug,omitempty" gtag:"_dbg,omitempty"`
	// If the current request has a referrer, it will be ignored at processing level
	// Example: 1
	// IgnoreReferrer *string `json:"ignore_referrer,omitempty" gtag:"ir,omitempty"`
	// Traffic Type
	// Example: 1
	// TrafficType *string `json:"traffic_type,omitempty" gtag:"tt,omitempty"`
	// Will be set to 1 is the current page has a linker and this last one is valid
	// Example: 1
	// IsGoogleLinkerValid *string `json:"is_google_linker_valid,omitempty" gtag:"_glv,omitempty"`

	// --- Event Parameters ---

	// Current Payload Name.
	// Example: page_view
	EventName *sesamy.EventName `json:"event_name,omitempty" gtag:"en,omitempty"`
	// It's the total engagement time in milliseconds since the last event. The engagement time is measured only when the current page is visible and active ( ie: the browser window/tab must be active and visible ), for this GA4 uses the window.events: focus, blur, pageshow, pagehide and the document:visibilitychange, these will determine when the timer starts and pauses
	// Example: 1234
	// EngagementTime *string `json:"engagement_time,omitempty" gtag:"_et,omitempty"`
	// Defines a parameter for the current Payload
	// Example: ep.page_type: checkout
	EventParameter map[string]string `json:"event_parameter,omitempty" gtag:"ep,omitempty"`
	// Defines a parameter for the current Payload
	// Example: epn.plays_count: 42
	EventParameterNumber map[string]string `json:"event_parameter_number,omitempty" gtag:"epn,omitempty"`
	// External Event
	// ExternalEvent *string `json:"external_event,omitempty" gtag:"_ee,omitempty"`

	// --- Session / User Related ---

	// Current User ID
	// Example: 1635691016
	UserID *string `json:"user_id,omitempty" gtag:"uid,omitempty"`
	// Current Firebase ID
	// Example: HASHSAH
	// FirebaseID *string `json:"firebase_id,omitempty" gtag:"_fid,omitempty"`
	// GA4 Session Id. This comes from the GA4 Cookie. It may be different for each Stream ID Configured on the site
	// Example: 1635691016
	SessionID *string `json:"session_id,omitempty" gtag:"sid,omitempty"`
	// Count of sessions recorded by GA4. This value increases by one each time a new session is detected ( when the session expires )
	// Example: 10
	// SessionCount *string `json:"session_count,omitempty" gtag:"sct,omitempty"`
	// If the current user is engaged in any way, this value will be 1
	// Example:
	// SessionEngagment *string `json:"session_engagement,omitempty" gtag:"seg,omitempty"`
	// Defines an user Propery for the current Measurement ID
	// Example: up.is_premium_user: yes
	UserProperty map[string]string `json:"user_property,omitempty" gtag:"up,omitempty"`
	// Defines an user Propery for the current Measurement ID
	// Example:
	UserPropertyNumber map[string]string `json:"user_property_number,omitempty" gtag:"upn,omitempty"`
	// If the "_ga_THYNGSTER" cookie is not set, the first event will have this value present. This will internally create a new "first_visit" event on GA4. If this event is also a conversion the value will be "2" if not, will be "1"
	// Example: 1|2
	// FirstVisit *string `json:"first_visit,omitempty" gtag:"_fv,omitempty"`
	// If the "_ga_THYNGSTER" cookie last session time value is older than 1800 seconds, the current event will have this value present. This will internally create a new "session_start" event on GA4. If this event is also a conversion the value will be "2" if not, will be "1"
	// Example: 1|2
	// SessionStart *string `json:"session_start,omitempty" gtag:"_ss,omitempty"`
	// This seems to be related to the ServerSide hits, it's 0 if the FPLC Cookie is not present and to the current value if it's coming from a Cross Domain linker
	// Example: bVhVicbfiSXaGNxeawKaPlDQc9QXPD6bKcsn36Elden6wZNb7Q5X1iXlkTVP5iP3H3y76cgM3UIgHCaRsYfPoyLGlbiIYMPRjvnUU7KWbdWLagodzxjrlPnvaRZJkw
	// FirstPartyLinkerCookie *string `json:"first_party_linker_cookie,omitempty" gtag:"_fplc,omitempty"`
	// If the current user has a GA4 session cookie, but not a GA (_ga) client id cookie, this parameter will be added to the hit
	// Example: 1
	// NewSessionID *string `json:"new_session_id,omitempty" gtag:"_nsi,omitempty"`
	// You may find this parameter if using some vendor plugin o platform ( ie: using shopify integration or a prestashop plugin )
	// Example: jdhsd87
	// GoogleDeveloperID *string `json:"google_developer_id,omitempty" gtag:"gdid,omitempty"`

	// --- Uncategorized / Missing Info ---

	// Example: 1
	// GTMUp *string `json:"gtmup,omitempty" gtag:"gtm_up,omitempty"`
	// Documented values, 1,2,3: Not sure when it's added.
	// EuropeanConsentModeEnabledID *string `json:"european_consent_mode_enabled_id,omitempty" gtag:"_ecid,omitempty"`
	// Example:
	// UEI *string `json:"uei,omitempty" gtag:"_uei,omitempty"`
	// It's set when a Google Join is created/imported. Google Signals
	// Example: 1
	// CreateGoogleJoin *string `json:"create_google_join,omitempty" gtag:"_gaz,omitempty"`
	// Example: Redact Device Info. Need Investigation about functionality
	// RedactDeviceInfo *string `json:"redact_device_info,omitempty" gtag:"_rdi,omitempty"`
	// Geo Granularity. Need Investigation about functionality
	// GeoGranularity *string `json:"geo_granularity,omitempty" gtag:"_geo,omitempty"`
	// Sent on sites that implement the US Privacy User Signal Mechanism, sent if window.__uspapi is present and returning a value.
	// Example: 1YNY
	// USPrivacySignal *string `json:"usprivacy_signal,omitempty" gtag:"us_privacy,omitempty"`
	// Sent on sites that implements IAB GDPR-Transparency-and-Consent-Framework( TCFv2 ) Mechanism. sent if window.__tcfapi is present and returning a valid value.
	// Example: 1
	// GDPR *string `json:"gdpr,omitempty" gtag:"gdpr,omitempty"`
	// Sent on sites that implements IAB GDPR-Transparency-and-Consent-Framework( TCFv2 ) Mechanism. sent if window.__tcfapi is present and returning a valid value.
	// Example: CPfPdAAPfPdAAAHABBENCgCsAP_AAAAAAAAAI_tf_X__b3_j-_5___t0eY1f9_7__-0zjhfdl-8N3f_X_L8X_2M7vF36tq4KuR4Eu3LBIQdlHOHcTUmw6okVrzPsbk2cr7NKJ7PEmnMbeydYGH9_n1_z-ZKY7_____77__-____3_____-_f___5_3____f_V__97fn9_____9_P___9v__9__________3___gAAAJJQAYAAgj-GgAwABBH8VABgACCP5SADAAEEfx0AGAAII_kIAMAAQR_CQAYAAgj-IgAwABBH8ZABgACCP4A.f_gAAAAAAAAA
	// GDPRConsent *string `json:"gdprconsent,omitempty" gtag:"gdpr_consent,omitempty"`
	// Example: sypham
	NonPersonalizedAds *string `json:"non_personalized_ads,omitempty" gtag:"npa,omitempty"`
	// Example: 1
	// ARE *string `json:"are,omitempty" gtag:"are,omitempty"`
	// PrivacySandboxCookieDeprecationLabel *string `json:"privacy_sandbox_cookie_deprecation_label,omitempty" gtag:"pscdl,omitempty"`
	// A timestamp measuring the difference between the moment this parameter gets populated and the moment the navigation started on that particular page.
	// TFD *string `json:"tfd,omitempty" gtag:"tfd,omitempty"`
	SST *SST `json:"sst,omitempty" gtag:"sst,omitempty"`
	// PAE *string `json:"pae,omitempty" gtag:"pae,omitempty"`

	// --- Unresolved ---

	Remain map[string]any `json:"-" gtag:"-,omitempy,remain"`
}
