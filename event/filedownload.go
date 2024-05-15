package event

import (
	"github.com/foomo/sesamy-go"
	"github.com/foomo/sesamy-go/event/params"
)

type FileDownload sesamy.Event[params.FileDownload]

func NewFileDownload(p params.FileDownload) FileDownload {
	return FileDownload(sesamy.NewEvent(sesamy.EventNameFileDownload, p))
}
