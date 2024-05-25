package event

import (
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type FileDownload sesamy.Event[params.FileDownload]

func NewFileDownload(p params.FileDownload) sesamy.Event[params.FileDownload] {
	return sesamy.NewEvent(sesamy.EventNameFileDownload, p)
}
