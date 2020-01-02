package main

import (
	"context"
	"fmt"
	"github.com/rs/cors"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/evscott/aedibus-api/dal"
	"github.com/evscott/aedibus-api/shared/logger"
	"github.com/google/go-github/github"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"golang.org/x/oauth2"
)

type Specifications struct {
	SrvPort           string `default:"8080"`
	ReadWriteTimeOut  string `default:"10"`
	HostIP            string `default:"127.0.0.1"`
	GithubAccessToken string `default:"no token provided"`
}

type dbSpecs struct {
	Migrations string `default:"file:///app/migrations"`
	Host       string `default:"db"`
	Port       string `default:"5432"`
	User       string `default:"user"`
	Password   string `default:"password"`
	Name       string `default:"dev"`
}

type Config struct {
	Spec         *Specifications
	Router       *mux.Router
	Server       *http.Server
	GithubClient *github.Client
	Logger       *logger.StandardLogger
	DAL          *dal.DAL
}

func Init(ctx context.Context, router *mux.Router) *Config {
	// Root log
	log := logger.NewLogger()

	/*****  Setup z3-12c-api specifications *****/
	spec := Specifications{}

	if err := envconfig.Process("Z3", &spec); err != nil {
		log.ConfigError(err)
	}
	// Get host IP
	if ipAddr, err := net.InterfaceAddrs(); err != nil {
		log.ConfigError(err)
	} else {
		spec.HostIP = strings.Split(ipAddr[0].String(), "/")[0]
	}

	/*****  Initialize Config *****/
	config := &Config{
		Spec:   &spec,
		Router: router,
		Server: &http.Server{
			Handler:      cors.Default().Handler(router),
			Addr:         fmt.Sprintf(":%s", spec.SrvPort),
			ReadTimeout:  time.Second * 10,
			WriteTimeout: time.Second * 10,
		},
		Logger: log,
	}

	/***** Setup Github client *****/
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: config.Spec.GithubAccessToken})
	tc := oauth2.NewClient(ctx, ts)
	githubClient := github.NewClient(tc)
	config.GithubClient = githubClient

	time.Sleep(time.Second * 2) // Snooze until database is spun up

	/***** Setup Postgres client & Database Access Layer (DAL) *****/
	db := dbSpecs{}
	if err := envconfig.Process("DB", &db); err != nil {
		log.ConfigError(err)
	}
	config.DAL = dal.Init(config.Logger, db.Host, db.Port, db.User, db.Password, db.Name, db.Migrations)

	return config
}
