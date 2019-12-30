package Routes

import (
	"github.com/evscott/aedibus-api/dal"
	"github.com/evscott/aedibus-api/router/handlers"
	"github.com/evscott/aedibus-api/shared/logger"
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
	c.level1()
	c.level2()
	c.level3()
}

func (c *Config) level1() {
}

func (c *Config) level2() {
	c.Router.HandleFunc(Path(Readme), c.Handlers.GetReadme).Methods(GET)
	c.Router.HandleFunc(Path(Submit, Assignment), c.Handlers.SubmitAssignment).Methods(POST)
	c.Router.HandleFunc(Path(Assignments), c.Handlers.GetAssignments).Methods(GET)
	c.Router.HandleFunc(Path(Submission, Feedback, File), c.Handlers.GetFeedbackOnSubmissionFile).Methods(GET)
	c.Router.HandleFunc(Path(File, Contents), c.Handlers.GetFileContents).Methods(GET)
	c.Router.HandleFunc(Path(Submission), c.Handlers.GetSubmissionResults).Methods(GET)
	c.Router.HandleFunc(Path(Submission, Feedback, File), c.Handlers.GetFeedbackOnSubmissionFile).Methods(GET)
}

func (c *Config) level3() {
	c.Router.HandleFunc(Path(Dropbox, File), c.Handlers.CreateDropboxFile).Methods(POST)
	c.Router.HandleFunc(Path(Assignment), c.Handlers.CreateAssignment).Methods(POST)
	c.Router.HandleFunc(Path(Assignment, File), c.Handlers.CreateAssignmentFile).Methods(POST)
	c.Router.HandleFunc(Path(Assignment, File), c.Handlers.UpdateAssignmentFile).Methods(PUT)
	c.Router.HandleFunc(Path(Dropbox), c.Handlers.CreateDropbox).Methods(POST)
	c.Router.HandleFunc(Path(Submission, Feedback), c.Handlers.LeaveCommentOnSubmission).Methods(POST)
}
