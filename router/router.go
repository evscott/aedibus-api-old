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
	c.Router.HandleFunc(Path(Github, File), c.Handlers.UploadFile).Methods(POST)
	// Update File
	c.Router.HandleFunc(Path(Github, File), c.Handlers.UpdateFile).Methods(PUT)
	// Get file
	c.Router.HandleFunc(Path(File), c.Handlers.GetFile).Methods(GET)
	// Get Readme
	c.Router.HandleFunc(Path(Readme), c.Handlers.GetReadme).Methods(GET)
}

func (c *Config) handleStudentRoutes() {
	// Create Submission
	c.Router.HandleFunc(Path(Github, Branch), c.Handlers.CreateSubmission).Methods(POST)
	// Submit Readme
	c.Router.HandleFunc(Path(Github, PullRequest), c.Handlers.SubmitAssignment).Methods(POST)
}

func (c *Config) handleInstructorRoutes() {
	// Create Readme
	c.Router.HandleFunc(Path(Github, Repository), c.Handlers.CreateAssignment).Methods(POST)
	// Create Comment on Readme
	c.Router.HandleFunc(Path(Github, PullRequest, Comment), c.Handlers.CreateComment).Methods(POST)
}
