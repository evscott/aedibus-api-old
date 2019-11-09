package Routes

import (
	"github.com/bndr/gojenkins"
	"github.com/evscott/z3-e2c-api/routes/handlers"
	"github.com/google/go-github/github"
	"github.com/gorilla/mux"
)

type Config struct {
	Router       *mux.Router
	GithubRoutes *handlers.Config
}

func GetRoutes(router *mux.Router, jenkins *gojenkins.Jenkins, github *github.Client) *Config {
	c := &Config{
		Router:       router,
		GithubRoutes: &handlers.Config{GAL: github},
	}

	c.handleGithubRoutes()

	return c
}

func (c *Config) handleGithubRoutes() {
	c.Router.HandleFunc(Path(Github, Repository), c.GithubRoutes.CreateRepository).Methods(POST)
	c.Router.HandleFunc(Path(Github, Branch), c.GithubRoutes.CreateReference).Methods(POST)
	c.Router.HandleFunc(Path(Github, File), c.GithubRoutes.UploadFile).Methods(POST)
	c.Router.HandleFunc(Path(Github, File), c.GithubRoutes.UpdateFile).Methods(PUT)
}
