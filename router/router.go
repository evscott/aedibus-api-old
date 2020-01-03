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
	c.Router.HandleFunc(Path(Assignments), c.Handlers.GetAssignments).Methods(GET)
}

func (c *Config) level3() {
	c.Router.HandleFunc(Path(Assignment), c.Handlers.CreateAssignment).Methods(POST)
	c.Router.HandleFunc(Path(Assignment), c.Handlers.DeleteAssignment).Methods(DELETE)
}
