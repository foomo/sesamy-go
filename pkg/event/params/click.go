package params

// Click holds the parameters of the enhanced measurement click event,
// sent for outbound link clicks. See https://support.google.com/analytics/answer/9216061.
type Click struct {
	LinkClasses string `json:"link_classes,omitempty"`
	LinkDomain  string `json:"link_domain,omitempty"`
	LinkID      string `json:"link_id,omitempty"`
	LinkURL     string `json:"link_url,omitempty"`
	Outbound    bool   `json:"outbound,omitempty"`
}
