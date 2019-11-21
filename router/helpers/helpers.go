package helpers

import (
	"github.com/evscott/z3-e2c-api/dal"
	"github.com/evscott/z3-e2c-api/shared/logger"
	"github.com/google/go-github/github"
)

type Config struct {
	DAL    *dal.DAL
	GAL    *github.Client
	Logger *logger.StandardLogger
}

func Init(dal *dal.DAL, gal *github.Client, logger *logger.StandardLogger) *Config {
	return &Config{
		DAL:    dal,
		GAL:    gal,
		Logger: logger,
	}
}
