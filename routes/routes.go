package Routesroutes

import (
	"github.com/bndr/gojenkins"
	"github.com/evscott/z3-e2c-api/routes/github-routes"
	consts "github.com/evscott/z3-e2c-api/shared"
	"github.com/google/go-github/github"
	"github.com/gorilla/mux"
)

type Config struct {
	Router       *mux.Router
	GithubRoutes *github_routes.Config
}

func GetRoutes(router *mux.Router, jenkins *gojenkins.Jenkins, github *github.Client) *Config {
	c := &Config{
		Router:       router,
		GithubRoutes: &github_routes.Config{GAL: github},
	}

	c.handleGithubRoutes()

	return c
}

func (c *Config) handleGithubRoutes() {
	c.Router.HandleFunc(consts.GITHUB, c.GithubRoutes.GetInfo).Methods(consts.GET)
	c.Router.HandleFunc(consts.GITHUB, c.GithubRoutes.Test).Methods(consts.POST)
	c.Router.HandleFunc(consts.GITHUB+consts.REPO, c.GithubRoutes.CreateRepository).Methods(consts.POST)
	c.Router.HandleFunc(consts.GITHUB+consts.BRANCH, c.GithubRoutes.CreateRef).Methods(consts.POST)
}
