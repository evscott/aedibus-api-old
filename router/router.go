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

func Init(router *mux.Router, dal *dal.DAL, github *github.Client, logger *logger.StandardLogger) {
	c := &Config{
		Router:   router,
		Handlers: handlers.Init(dal, github, logger),
	}

	c.handleGithubRoutes()
}

func (c *Config) handleGithubRoutes() {
	c.handleGeneralRoutes()
	c.handleStudentRoutes()
	c.handleInstructorRoutes()
}

func (c *Config) handleGeneralRoutes() {
	// Upload File
	c.Router.HandleFunc(Path(Github, File), c.Handlers.UploadAssignment).Methods(POST)
	// Update File
	c.Router.HandleFunc(Path(Github, File), c.Handlers.UpdateAssignment).Methods(PUT)
}

func (c *Config) handleStudentRoutes() {
	// Create Submission
	c.Router.HandleFunc(Path(Github, Branch), c.Handlers.CreateSubmission).Methods(POST)
	// Submit Assignment
	c.Router.HandleFunc(Path(Github, PullRequest), c.Handlers.SubmitAssignment).Methods(POST)
}

func (c *Config) handleInstructorRoutes() {
	// Create Assignment
	c.Router.HandleFunc(Path(Github, Repository), c.Handlers.CreateAssignment).Methods(POST)
	// Create Comment on Assignment
	c.Router.HandleFunc(Path(Github, PullRequest, Comment), c.Handlers.CreateComment).Methods(POST)
}
