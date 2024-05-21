package params

type Click struct {
	LinkClasses string `json:"link_classes,omitempty"`
	LinkDomain  string `json:"link_domain,omitempty"`
	LinkID      string `json:"link_id,omitempty"`
	LinkURL     string `json:"link_url,omitempty"`
	Outbound    bool   `json:"outbound,omitempty"`
}
