package handlers

import (
	"github.com/evscott/aedibus-api/dal"
	"github.com/evscott/aedibus-api/router/helpers"
	"github.com/evscott/aedibus-api/shared/logger"
	"github.com/google/go-github/github"
)

type Config struct {
	logger  *logger.StandardLogger
	helpers *helpers.Config
}

func Init(dal *dal.DAL, gal *github.Client, logger *logger.StandardLogger) *Config {
	return &Config{
		logger:  logger,
		helpers: helpers.Init(dal, gal),
	}
}
