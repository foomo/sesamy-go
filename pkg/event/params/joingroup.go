package params

// JoinGroup https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference/events#join_group
type JoinGroup struct {
	GroupID string `json:"group_id,omitempty"`
}
