package params

// UnlockAchievement https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference/events#unlock_achievement
type UnlockAchievement struct {
	ArchievementID string `json:"achievement_id"`
}
