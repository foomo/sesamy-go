package http

import (
	"net/http"

	"github.com/foomo/sesamy-go/pkg/sesamy"
	"go.uber.org/zap"
)

// EventHandler processes a single event of an incoming request. Handlers may
// modify event in place. A non-nil error aborts processing of the request.
type EventHandler func(l *zap.Logger, r *http.Request, event *sesamy.Event[any]) error
