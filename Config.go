package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/evscott/z3-e2c-api/shared/logger"
	"github.com/google/go-github/github"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"golang.org/x/oauth2"
)

type Specifications struct {
	SrvPort           string `default:"7070"`
	ReadWriteTimeOut  string `default:"10"`
	HostIP            string `default:"127.0.0.1"`
	GithubAccessToken string
}

type Config struct {
	Spec         *Specifications
	Router       *mux.Router
	Server       *http.Server
	GithubClient *github.Client
	Logger       *logger.StandardLogger
}

func GetConfig(ctx context.Context, router *mux.Router) *Config {
	// Setup logger
	log := logger.NewLogger()

	/*****  Setup z3-12c-api specifications *****/
	spec := Specifications{}
	// Load environment variables from .env if found
	err := godotenv.Load()
	if err != nil {
		log.ConfigError(err)
	}
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
			Handler:      router,
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

	/***** Run Postgres migrations client *****/

	return config
}
