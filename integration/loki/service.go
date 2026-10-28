package loki

import (
	"context"
)

// Service runs a [Loki] as a named service with Start and Close.
type Service struct {
	loki *Loki
}

// NewService returns a Service that runs loki.
func NewService(loki *Loki) *Service {
	return &Service{loki: loki}
}

// Name returns "loki".
func (s *Service) Name() string {
	return "loki"
}

// Start runs [Loki.Start]. It blocks until ctx is done or Close is called.
func (s *Service) Start(ctx context.Context) error {
	s.loki.Start(ctx)
	return nil
}

// Close stops the Loki, see [Loki.Stop].
func (s *Service) Close(ctx context.Context) error {
	s.loki.Stop()
	return nil
}
