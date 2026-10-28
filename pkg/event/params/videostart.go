package params

// VideoStart holds the parameters of the enhanced measurement video_start
// event. See https://support.google.com/analytics/answer/9216061.
type VideoStart struct {
	VideoCurrentTime int64  `json:"video_current_time,omitempty"`
	VideoDuration    int64  `json:"video_duration,omitempty"`
	VideoPercent     int64  `json:"video_percent,omitempty"`
	VideoProvider    string `json:"video_provider,omitempty"`
	VideoTitle       string `json:"video_title,omitempty"`
	VideoURL         string `json:"video_url,omitempty"`
	Visible          bool   `json:"visible,omitempty"`
}
