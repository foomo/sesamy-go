package params

// FileDownload holds the parameters of the enhanced measurement file_download
// event. See https://support.google.com/analytics/answer/9216061.
type FileDownload struct {
	FileExtension string `json:"file_extension,omitempty"`
	FileName      string `json:"file_name,omitempty"`
	LinkClasses   string `json:"link_classes,omitempty"`
	LinkID        string `json:"link_id,omitempty"`
	LinkText      string `json:"link_text,omitempty"`
	LinkURL       string `json:"link_url,omitempty"`
}
