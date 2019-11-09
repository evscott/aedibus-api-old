package Routesroutes

import (
	"github.com/bndr/gojenkins"
	"github.com/evscott/z3-12c-api/Routes/GithubRoutes"
	"github.com/evscott/z3-12c-api/Routes/JenkinsRoutes"
	"github.com/google/go-github/github"
	"github.com/gorilla/mux"
)

type Config struct {
	Router        *mux.Router
	JenkinsRoutes *JenkinsRoutes.Config
	GithubRoutes  *GithubRoutes.Config
}

func GetRoutes(router *mux.Router, jenkins *gojenkins.Jenkins, github *github.Client) *Config {
	c := &Config{
		Router:        router,
		JenkinsRoutes: &JenkinsRoutes.Config{JAL: jenkins},
		GithubRoutes:  &GithubRoutes.Config{GAL: github},
	}

	_ = c.Router.HandleFunc("/jenkins", c.JenkinsRoutes.Test)
	_ = c.Router.HandleFunc("/github", c.GithubRoutes.Test)

	return c
}
