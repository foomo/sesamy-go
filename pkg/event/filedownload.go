package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

// FileDownload is the GA4 file_download event.
type FileDownload sesamy.Event[params.FileDownload]

// NewFileDownload returns a file_download event with the parameters p.
func NewFileDownload(p params.FileDownload) sesamy.Event[params.FileDownload] {
	return sesamy.NewEvent(sesamy.EventNameFileDownload, p)
}
