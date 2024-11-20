package http

import (
	"net/http"

	"github.com/foomo/sesamy-go/pkg/sesamy"
)

type EventHandler func(r *http.Request, event *sesamy.Event[any]) error
