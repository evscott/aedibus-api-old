package Routes

import (
	"github.com/evscott/z3-e2c-api/routes/handlers"
	"github.com/evscott/z3-e2c-api/shared/logger"
	"github.com/google/go-github/github"
	"github.com/gorilla/mux"
)

type Config struct {
	Router       *mux.Router
	GithubRoutes *handlers.Config
}

func Init(router *mux.Router, github *github.Client, logger *logger.StandardLogger) {
	c := &Config{
		Router:       router,
		GithubRoutes: &handlers.Config{GAL: github, Logger: logger},
	}

	c.handleGithubRoutes()
}

func (c *Config) handleGithubRoutes() {
	c.handleGeneralRoutes()
	c.handleInstructorRoutes()
}

func (c *Config) handleGeneralRoutes() {
	// Upload File
	c.Router.HandleFunc(Path(Github, File), c.GithubRoutes.UploadFile).Methods(POST)
	// Update File
	c.Router.HandleFunc(Path(Github, File), c.GithubRoutes.UpdateFile).Methods(PUT)
	// Create Pull Request
	c.Router.HandleFunc(Path(Github, PullRequest), c.GithubRoutes.CreatePullRequest).Methods(POST)
}

func (c *Config) handleInstructorRoutes() {
	// Create Repository
	c.Router.HandleFunc(Path(Github, Repository), c.GithubRoutes.CreateRepository).Methods(POST)
	// Create Branch
	c.Router.HandleFunc(Path(Github, Branch), c.GithubRoutes.CreateBranch).Methods(POST)
	// Create Comment on Pull Request
	c.Router.HandleFunc(Path(Github, PullRequest, Comment), c.GithubRoutes.CreateComment).Methods(POST)
}
