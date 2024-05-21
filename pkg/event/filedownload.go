package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	sesamy2 "github.com/foomo/sesamy-go/pkg/sesamy"
)

type FileDownload sesamy2.Event[params.FileDownload]

func NewFileDownload(p params.FileDownload) FileDownload {
	return FileDownload(sesamy2.NewEvent(sesamy2.EventNameFileDownload, p))
}
