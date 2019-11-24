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

	c.initRoutes()
}

func (c *Config) initRoutes() {
	c.generalRoutes()
	c.studentRoutes()
	c.instructorRoutes()
}

func (c *Config) generalRoutes() {
	c.Router.HandleFunc(Path(Readme), c.Handlers.GetReadme).Methods(GET)
}

func (c *Config) studentRoutes() {
	c.Router.HandleFunc(Path(Dropbox, File), c.Handlers.CreateDropboxFile).Methods(POST)
	c.Router.HandleFunc(Path(Submit, Assignment), c.Handlers.SubmitAssignment).Methods(POST)
}

func (c *Config) instructorRoutes() {
	c.Router.HandleFunc(Path(Assignment), c.Handlers.CreateAssignment).Methods(POST)
	c.Router.HandleFunc(Path(Assignment, File), c.Handlers.CreateAssignmentFile).Methods(POST)
	c.Router.HandleFunc(Path(Assignment, File), c.Handlers.UpdateAssignmentFile).Methods(PUT)
	c.Router.HandleFunc(Path(File, Contents), c.Handlers.GetFileContents).Methods(GET)
	c.Router.HandleFunc(Path(Dropbox), c.Handlers.CreateDropbox).Methods(POST)
	c.Router.HandleFunc(Path(Submission), c.Handlers.GetSubmissionResults).Methods(GET)
	c.Router.HandleFunc(Path(Submission, Feedback), c.Handlers.LeaveFeedbackOnSubmission).Methods(POST)
}
