package handlers

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/evscott/z3-e2c-api/models"
	consts "github.com/evscott/z3-e2c-api/shared/constants"
	"github.com/evscott/z3-e2c-api/shared/utils"
	"github.com/google/go-github/github"
)

type Config struct {
	GAL *github.Client
}

//  TODO
func (c *Config) CreateComment(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Unpack create comment request
	req := &models.ReqCreateComment{}
	UnmarshalRequest(req, w, r)

	comment := github.PullRequestComment{
		Path:     req.Path,
		Body:     req.Body,
		Position: req.Position,
		CommitID: req.CommitID,
	}

	if res, _, err := c.GAL.PullRequests.CreateComment(ctx, consts.Z3E2C, *req.RepoName, 1, &comment); err != nil {
		w.WriteHeader(Status(InternalServerError))
		log.Fatal(err)
	} else {
		MarshalResponse(res, w)
	}
}

func (c *Config) CreatePullRequest(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Unpack create repository request
	req := &models.ReqCreatePR{}
	UnmarshalRequest(req, w, r)

	pullRequest := github.NewPullRequest{
		Title:               req.Title,
		Head:                req.Head,
		Base:                utils.String(consts.MASTER),
		Body:                req.Body,
		Issue:               nil,
		MaintainerCanModify: utils.Bool(true),
	}

	if res, _, err := c.GAL.PullRequests.Create(ctx, consts.Z3E2C, *req.RepoName, &pullRequest); err != nil {
		w.WriteHeader(Status(InternalServerError))
		log.Fatal(err)
	} else {
		MarshalResponse(res, w)
	}
}

// UpdateFile TODO
func (c *Config) UpdateFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Unpack file metadata
	repo := r.FormValue("repo")
	branch := r.FormValue("branch")
	fileName := r.FormValue("fileName")

	// Unpack request to update file
	file, header, err := r.FormFile(fileName)
	if err != nil {
		w.WriteHeader(Status(InternalServerError))
		log.Fatal(err)
	} else {
		name := strings.Split(header.Filename, ".")
		fmt.Printf("Received file: %s\n", name[0])
		defer file.Close()
	}

	// Unpack file into byte array
	buffer := bytes.Buffer{}
	if _, err := io.Copy(&buffer, file); err != nil {
		w.WriteHeader(Status(InternalServerError))
		log.Fatal(err)
	}
	contents := buffer.Bytes()
	buffer.Reset()

	// Get blob sha of file from Github to be used as target of update
	var sha string
	getOptions := github.RepositoryContentGetOptions{Ref: fmt.Sprintf("heads/%s", branch)}
	if contents, _, res, err := c.GAL.Repositories.GetContents(ctx, consts.Z3E2C, repo, fileName, &getOptions); err != nil {
		w.WriteHeader(Status(InternalServerError))
		log.Fatal(err)
	} else {
		fmt.Printf("Got sha for file %s %v\n", fileName, res)
		sha = *contents.SHA
	}

	// Upload file to Github
	fileOptions := github.RepositoryContentFileOptions{
		Message: utils.String("Updating file"), // TODO
		Content: contents,
		Branch:  &branch,
		SHA:     &sha,
	}
	if res, _, err := c.GAL.Repositories.UpdateFile(ctx, consts.Z3E2C, repo, fileName, &fileOptions); err != nil {
		w.WriteHeader(Status(InternalServerError))
		log.Fatal(err)
	} else {
		MarshalResponse(res, w)
	}
}

// UploadFile TODO
func (c *Config) UploadFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Unpack file metadata
	repo := r.FormValue("repo")
	branch := r.FormValue("branch")
	fileName := r.FormValue("fileName")

	// Unpack request to upload file
	file, header, err := r.FormFile(fileName)
	if err != nil {
		w.WriteHeader(Status(InternalServerError))
		log.Fatal(err)
	} else {
		name := strings.Split(header.Filename, ".")
		fmt.Printf("Received file: %s\n", name[0])
		defer file.Close()
	}

	// Unpack file into byte array
	buffer := bytes.Buffer{}
	if _, err := io.Copy(&buffer, file); err != nil {
		w.WriteHeader(Status(InternalServerError))
		log.Fatal(err)
	}
	contents := buffer.Bytes()
	buffer.Reset()

	// Upload file to Github
	fileOptions := github.RepositoryContentFileOptions{
		Message: utils.String("Uploading file"),
		Content: contents,
		Branch:  &branch,
	}
	if res, _, err := c.GAL.Repositories.CreateFile(ctx, consts.Z3E2C, repo, fileName, &fileOptions); err != nil {
		w.WriteHeader(Status(InternalServerError))
		log.Fatal(err)
	} else {
		MarshalResponse(res, w)
	}
}

// CreateRepository TODO
func (c *Config) CreateRepository(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Unpack create repository request
	req := &models.ReqCreateRepo{}
	UnmarshalRequest(req, w, r)

	// Create repository
	repo := github.Repository{
		Name:          req.RepoName,
		DefaultBranch: utils.String(consts.MASTER),
	}
	if res, _, err := c.GAL.Repositories.Create(ctx, consts.Z3E2C, &repo); err != nil {
		w.WriteHeader(Status(InternalServerError))
		log.Fatal(err)
	} else {
		MarshalResponse(res, w)
	}
}

// CreateBranch TODO
func (c *Config) CreateBranch(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Unpack create reference request
	req := &models.ReqCreateBranch{}
	UnmarshalRequest(req, w, r)

	// Get MASTER reference
	masterRef, res, err := c.GAL.Git.GetRef(ctx, consts.Z3E2C, *req.RepoName, fmt.Sprintf("refs/heads/%s", consts.MASTER))
	if err != nil {
		w.WriteHeader(Status(InternalServerError))
		log.Fatal(err)
	} else {
		fmt.Printf("Got consts.MASTER reference: %v\n", res)
	}

	// Create branch
	reference := github.Reference{
		Ref: utils.String(fmt.Sprintf("refs/heads/%s", req.BranchName)),
		URL: utils.String(fmt.Sprintf("https://api.github.com/repos/%s/%s/git/refs/heads/%s", consts.Z3E2C, req.RepoName, req.BranchName)),
		Object: &github.GitObject{
			Type: utils.String("commit"),
			SHA:  masterRef.Object.SHA,
			URL:  utils.String(fmt.Sprintf("https://api.github.com/repos/%s/%s/git/commits/%s", consts.Z3E2C, req.RepoName, consts.MASTER)),
		},
	}
	if res, _, err := c.GAL.Git.CreateRef(ctx, consts.Z3E2C, *req.RepoName, &reference); err != nil {
		w.WriteHeader(Status(InternalServerError))
		log.Fatal(err)
	} else {
		MarshalResponse(res, w)
	}
}
