package cmd

import (
	"github.com/enescakir/emoji"
	"github.com/kuritka/golic/utils/guard"
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
	logger.Info().Msgf("%s command %s started",emoji.Tractor, r.service)
	err := r.service.Run()
	guard.FailOnError(err, "command %s failed", r.service)
}
