package params

type VideoComplete struct {
	VideoCurrentTime int64  `json:"video_current_time,omitempty"`
	VideoDuration    int64  `json:"video_duration,omitempty"`
	VideoPercent     int64  `json:"video_percent,omitempty"`
	VideoProvider    string `json:"video_provider,omitempty"`
	VideoTitle       string `json:"video_title,omitempty"`
	VideoURL         string `json:"video_url,omitempty"`
	Visible          bool   `json:"visible,omitempty"`
}
