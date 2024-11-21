package http

import (
	"net/http"

	"github.com/foomo/sesamy-go/pkg/sesamy"
	"go.uber.org/zap"
)

type EventHandler func(l *zap.Logger, r *http.Request, event *sesamy.Event[any]) error
