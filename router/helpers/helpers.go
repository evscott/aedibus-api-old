package helpers

import (
	"github.com/evscott/z3-e2c-api/dal"
	"github.com/evscott/z3-e2c-api/router/helpers/db"
	"github.com/evscott/z3-e2c-api/router/helpers/gh"
	"github.com/google/go-github/github"
)

type Config struct {
	DB db.Provider
	GH gh.Provider
}

func Init(dal *dal.DAL, gal *github.Client) *Config {
	return &Config{
		DB: db.Init(dal),
		GH: gh.Init(gal),
	}
}
