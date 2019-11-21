package Routes

import (
	"github.com/evscott/z3-e2c-api/dal"
	"github.com/evscott/z3-e2c-api/router/handlers"
	"github.com/evscott/z3-e2c-api/shared/logger"
	"github.com/google/go-github/github"
	"github.com/gorilla/mux"
)

type Config struct {
	Router   *mux.Router
	Handlers *handlers.Config
}

func Init(router *mux.Router, github *github.Client, dal *dal.DAL, logger *logger.StandardLogger) {
	c := &Config{
		Router:   router,
		Handlers: &handlers.Config{DAL: dal, GAL: github, Logger: logger},
	}

	c.handleGithubRoutes()
}

func (c *Config) handleGithubRoutes() {
	c.handleGeneralRoutes()
	c.handleInstructorRoutes()
}

func (c *Config) handleGeneralRoutes() {
	// Upload File
	c.Router.HandleFunc(Path(Github, File), c.Handlers.UploadAssignment).Methods(POST)
	// Update File
	c.Router.HandleFunc(Path(Github, File), c.Handlers.UpdateAssignment).Methods(PUT)
	// Create Pull Request
	c.Router.HandleFunc(Path(Github, PullRequest), c.Handlers.SubmitAssignment).Methods(POST)
}

func (c *Config) handleInstructorRoutes() {
	// Create Repository
	c.Router.HandleFunc(Path(Github, Repository), c.Handlers.CreateAssignment).Methods(POST)
	// Create Branch
	c.Router.HandleFunc(Path(Github, Branch), c.Handlers.CreateSubmission).Methods(POST)
	// Create Comment on Pull Request
	c.Router.HandleFunc(Path(Github, PullRequest, Comment), c.Handlers.CreateComment).Methods(POST)
}
