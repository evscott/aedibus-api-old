package github_routes

import (
	"context"
	"fmt"
	"github.com/evscott/z3-e2c-api/models"
	"log"
	"net/http"

	consts "github.com/evscott/z3-e2c-api/shared"
	"github.com/google/go-github/github"
)

type Config struct {
	GAL *github.Client
}

func (c *Config) GetInfo(w http.ResponseWriter, r *http.Request) {}

func (c *Config) Test(w http.ResponseWriter, r *http.Request) {
	req := &models.ReqCreateRef{}
	consts.ParseReqJsonBody(req, w, r)
	fmt.Printf("Body: %v\n", req)
	w.WriteHeader(200)
}

func (c *Config) CreateRepository(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	/***** Unpack create repository request *****/
	req := &models.ReqCreateRepo{}
	consts.ParseReqJsonBody(req, w, r)
	if req.Repo == "" {
		w.WriteHeader(400)
		log.Fatalf("Invalid request: %v\n", req)
	}

	/***** Create repository *****/
	repo := github.Repository{
		Name:          &req.Repo,
		DefaultBranch: String(consts.MASTER),
	}
	if _, res, err := c.GAL.Repositories.Create(ctx, consts.Z3E2C, &repo); err != nil {
		w.WriteHeader(500)
		log.Fatal(err)
	} else {
		fmt.Printf("Repository %s created %v\n", req.Repo, res)
	}

	w.WriteHeader(200)
}

func (c *Config) CreateReference(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	/***** Unpack create reference request *****/
	req := &models.ReqCreateRef{}
	consts.ParseReqJsonBody(req, w, r)
	if req.Repo == "" || req.Branch == "" {
		w.WriteHeader(400)
		log.Fatalf("Invalid request: %v\n", req)
	}

	/***** Get master reference *****/
	masterRef, res, err := c.GAL.Git.GetRef(ctx, consts.Z3E2C, req.Repo, fmt.Sprintf("refs/heads/%s", consts.MASTER))
	if err != nil {
		w.WriteHeader(500)
		log.Fatal(err)
	} else {
		fmt.Printf("Got master reference: %v\n", res)
	}

	/***** Create branch *****/
	reference := github.Reference{
		Ref: String(fmt.Sprintf("refs/heads/%s", req.Branch)),
		URL: String(fmt.Sprintf("https://api.github.com/repos/%s/%s/git/refs/heads/%s", consts.Z3E2C, req.Repo, req.Branch)),
		Object: &github.GitObject{
			Type: String("commit"),
			SHA:  masterRef.Object.SHA,
			URL:  String(fmt.Sprintf("https://api.github.com/repos/%s/%s/git/commits/%s", consts.Z3E2C, req.Repo, masterRef)),
		},
	}

	if _, res, err := c.GAL.Git.CreateRef(ctx, consts.Z3E2C, req.Repo, &reference); err != nil {
		w.WriteHeader(500)
		log.Fatal(err)
	} else {
		fmt.Printf("Reference %s/%s created: %v\n", req.Repo, req.Branch, res)
	}

	w.WriteHeader(200)
}

func String(s string) *string {
	return &s
}
