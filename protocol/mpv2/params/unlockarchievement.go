package params

// UnlockArchievement https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference/events#unlock_achievement
type UnlockArchievement struct {
	ArchievementName string `json:"archievement_name"`
}
