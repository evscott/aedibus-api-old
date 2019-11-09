package main

import (
	"context"
	"fmt"
	"github.com/bndr/gojenkins"
	"github.com/google/go-github/github"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"golang.org/x/oauth2"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

type Specifications struct {
	JenkinsHost       string `default:"http://localhost:8080"`
	JenkinsUsername   string `default:"admin"`
	JenkinsPassword   string `default:"admin"`
	SrvPort           string `default:"9090"`
	ReadWriteTimeOut  string `default:"10"`
	HostIP            string `default:"127.0.0.1"`
	GithubAccessToken string
}

type Config struct {
	Spec          *Specifications
	Router        *mux.Router
	Server        *http.Server
	JenkinsClient *gojenkins.Jenkins
	GithubClient  *github.Client
}

func GetConfig(ctx context.Context, router *mux.Router) *Config {
	/*****  Setup z3-12c-api specifications *****/
	spec := Specifications{}
	// Load environment variables from .env if found
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	if err := envconfig.Process("Z3", &spec); err != nil {
		log.Fatal(err)
	}
	// Get host IP
	if ipAddr, err := net.InterfaceAddrs(); err != nil {
		log.Fatal(err)
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
	}

	/***** Set up Jenkins client *****/
	jenkinsClient, _ := gojenkins.CreateJenkins(nil, spec.JenkinsHost, spec.JenkinsUsername, spec.JenkinsPassword).Init()
	if jenkinsClient != nil {
		nodes, err := jenkinsClient.GetAllNodes()
		if err != nil {
			log.Fatal(err)
		}
		// Check if all nodes are online
		for _, node := range nodes {
			if _, err := node.Poll(); err != nil {
				log.Fatal(err)
			}
			if isOnline, err := node.IsOnline(); err == nil {
				fmt.Printf("Node %s is online: %v\n", node.GetName(), isOnline)
			}
		}
		config.JenkinsClient = jenkinsClient
	}

	/***** Setup Github client *****/
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: config.Spec.GithubAccessToken})
	tc := oauth2.NewClient(ctx, ts)
	githubClient := github.NewClient(tc)
	config.GithubClient = githubClient

	return config
}
