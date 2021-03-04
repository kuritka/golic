package cmd

import (
	"gext/guard"
)

type Service interface {
	Run() error
	String() string
}

type ServiceRunner struct {
	service Service
}

func Command(service Service) *ServiceRunner {
	return &ServiceRunner{
		service,
	}
}

//Run service once and panics if service is broken
func (r *ServiceRunner) MustRun() {
	logger.Info().Msgf("service %s started", r.service)
	err := r.service.Run()
	guard.FailOnError(err, "service %s failed", r.service)
}
