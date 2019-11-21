package handlers

import (
	"github.com/evscott/z3-e2c-api/dal"
	"github.com/evscott/z3-e2c-api/router/handlers/helpers"
	"github.com/evscott/z3-e2c-api/shared/logger"
	"github.com/google/go-github/github"
)

type Config struct {
	DAL     *dal.DAL
	GAL     *github.Client
	Logger  *logger.StandardLogger
	helpers *helpers.Config
}

func Init(dal *dal.DAL, gal *github.Client, logger *logger.StandardLogger) *Config {
	return &Config{
		DAL:     dal,
		GAL:     gal,
		Logger:  logger,
		helpers: helpers.Init(dal, gal, logger),
	}
}
