package GithubRoutes

import (
	"context"
	"fmt"
	"log"
	"net/http"

	consts "github.com/evscott/z3-e2c-api/shared"
	"github.com/google/go-github/github"
)

type Config struct {
	GAL *github.Client
}

func (c *Config) GetInfo(w http.ResponseWriter, r *http.Request) {}

func (c *Config) CreateRepository(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	name := "test"
	defaultBranch := "master"

	repo := github.Repository{
		Name:          &name,
		DefaultBranch: &defaultBranch,
	}

	if repo, res, err := c.GAL.Repositories.Create(ctx, consts.Z3E2C, &repo); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Status: %s\nRepository created: %v\n", res, repo)
	}
}
