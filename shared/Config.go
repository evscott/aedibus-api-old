package shared

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"database/sql"
	"github.com/evscott/z3-e2c-api/shared/constants"
	"github.com/evscott/z3-e2c-api/shared/logger"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/go-github/github"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"golang.org/x/oauth2"
)

type Specifications struct {
	SrvPort           string `default:"8080"`
	ReadWriteTimeOut  string `default:"10"`
	HostIP            string `default:"127.0.0.1"`
	GithubAccessToken string `default:"no token provided"`

	DbMigrations string `default:"file:///app/migrations"`
	DbHost       string `default:"db"`
	DbUser       string `default:"user"`
	DbPassword   string `default:"password"`
	DbName       string `default:"dev"`
}

type Config struct {
	Spec         *Specifications
	Router       *mux.Router
	Server       *http.Server
	GithubClient *github.Client
	Logger       *logger.StandardLogger
	DbClient     *sql.DB
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

	time.Sleep(time.Second * 2) // Snooze until database is spun up

	/***** Setup database client *****/
	dbInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", config.Spec.DbHost, config.Spec.DbUser, config.Spec.DbPassword, config.Spec.DbName)
	if config.DbClient, err = sql.Open(constants.DB_DRIVER, dbInfo); err != nil {
		log.ConfigError(err)
	} else {
		log.Printf("Successfully connected to database: %s", config.Spec.DbName)
	}

	/***** Run database migrations *****/
	driver, err := postgres.WithInstance(config.DbClient, &postgres.Config{})
	if err != nil {
		log.ConfigError(err)
	}
	m, err := migrate.NewWithDatabaseInstance(config.Spec.DbMigrations, config.Spec.DbName, driver)
	if err != nil {
		log.ConfigError(err)
	}
	if err := m.Up(); err != nil {
		log.ConfigError(err)
	} else {
		log.Printf("Successfully migrated database")
	}

	return config
}
