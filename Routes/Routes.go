package Routesroutes

import (
	"github.com/bndr/gojenkins"
	"github.com/evscott/z3-e2c-api/Routes/GithubRoutes"
	consts "github.com/evscott/z3-e2c-api/shared"
	"github.com/google/go-github/github"
	"github.com/gorilla/mux"
)

type Config struct {
	Router       *mux.Router
	GithubRoutes *GithubRoutes.Config
}

func GetRoutes(router *mux.Router, jenkins *gojenkins.Jenkins, github *github.Client) *Config {
	c := &Config{
		Router:       router,
		GithubRoutes: &GithubRoutes.Config{GAL: github},
	}

	c.handleGithubRoutes()

	return c
}

func (c *Config) handleGithubRoutes() {
	c.Router.HandleFunc(consts.Github, c.GithubRoutes.GetInfo).Methods(consts.GET)
	c.Router.HandleFunc(consts.Github+consts.Create+consts.Repo, c.GithubRoutes.CreateRepository).Methods(consts.POST)
}
