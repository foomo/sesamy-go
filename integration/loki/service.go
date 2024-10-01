package loki

import (
	"context"
)

type Service struct {
	loki *Loki
}

func NewService(loki *Loki) *Service {
	return &Service{loki: loki}
}

func (s *Service) Name() string {
	return "loki"
}

// Start pulls lines out of the channel and sends them to Loki
func (s *Service) Start(ctx context.Context) error {
	s.loki.Start(ctx)
	return nil
}

// Close will cancel any ongoing requests and stop the goroutine listening for requests
func (s *Service) Close(ctx context.Context) error {
	s.loki.Stop()
	return nil
}
