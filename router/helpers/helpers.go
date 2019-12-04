package helpers

import (
	"github.com/evscott/aedibus-api/dal"
	"github.com/evscott/aedibus-api/router/helpers/db"
	"github.com/evscott/aedibus-api/router/helpers/gh"
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
