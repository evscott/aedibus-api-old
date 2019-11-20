package handlers

import (
	"bytes"
	"context"
	"fmt"
	"github.com/evscott/z3-e2c-api/dal"
	http2 "github.com/evscott/z3-e2c-api/shared/http"
	"github.com/evscott/z3-e2c-api/shared/marsh"
	"io"
	"net/http"
	"strings"

	"github.com/evscott/z3-e2c-api/models"
	consts "github.com/evscott/z3-e2c-api/shared/constants"
	"github.com/evscott/z3-e2c-api/shared/logger"
	"github.com/evscott/z3-e2c-api/shared/utils"
	"github.com/google/go-github/github"
)

type Config struct {
	DAL    *dal.DAL
	GAL    *github.Client
	Logger *logger.StandardLogger
}

//  TODO
//
//
func (c *Config) CreateComment(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Unmarshal create comment request
	req := &models.ReqCreateComment{}
	marsh.UnmarshalRequest(req, w, r)

	// Create comment
	comment := github.PullRequestComment{
		Path:     req.Path,
		Body:     req.Body,
		Position: req.Position,
		CommitID: req.CommitID,
	}
	if res, _, err := c.GAL.PullRequests.CreateComment(ctx, consts.Z3E2C, *req.RepoName, 1, &comment); err != nil {
		w.WriteHeader(http2.Status(http2.InternalServerError))
		c.Logger.GalError(err)
	} else { // Success
		marsh.MarshalResponse(res, w)
	}
}

// CreatePullRequest TODO
//
//
func (c *Config) CreatePullRequest(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Unmarshal create repository request
	req := &models.ReqCreatePR{}
	marsh.UnmarshalRequest(req, w, r)

	// Create pull request
	pullRequest := github.NewPullRequest{
		Title:               req.Title,
		Head:                req.Head,
		Base:                utils.String(consts.MASTER),
		Body:                req.Body,
		Issue:               nil,
		MaintainerCanModify: utils.Bool(true),
	}
	if res, _, err := c.GAL.PullRequests.Create(ctx, consts.Z3E2C, *req.RepoName, &pullRequest); err != nil {
		w.WriteHeader(http2.Status(http2.InternalServerError))
		c.Logger.GalError(err)
	} else { // Success
		marsh.MarshalResponse(res, w)
	}
}

// UpdateFigithub-httple TODO
//
//
func (c *Config) UpdateFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Extract values from requests forms
	repo := r.FormValue("repo")
	branch := r.FormValue("branch")
	fileName := r.FormValue("fileName")

	// Extract file from request form
	file, header, err := r.FormFile(fileName)
	if err != nil {
		w.WriteHeader(http2.Status(http2.InternalServerError))
		c.Logger.GalError(err)
	} else {
		name := strings.Split(header.Filename, ".")
		fmt.Printf("Received file: %s\n", name[0])
		defer file.Close()
	}

	// Read file contents through buffer
	buffer := bytes.Buffer{}
	if _, err := io.Copy(&buffer, file); err != nil {
		w.WriteHeader(http2.Status(http2.InternalServerError))
		c.Logger.GalError(err)
	}
	contents := buffer.Bytes()
	buffer.Reset()

	// Get blob sha of file from Github to be used as target of update
	var sha string
	getOptions := github.RepositoryContentGetOptions{Ref: fmt.Sprintf("heads/%s", branch)}
	if contents, _, res, err := c.GAL.Repositories.GetContents(ctx, consts.Z3E2C, repo, fileName, &getOptions); err != nil {
		w.WriteHeader(http2.Status(http2.InternalServerError))
		c.Logger.GalError(err)
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
		w.WriteHeader(http2.Status(http2.InternalServerError))
		c.Logger.GalError(err)
	} else { // Success
		marsh.MarshalResponse(res, w)
	}
}

// UploadFile TODO
//
//
func (c *Config) UploadFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Extract values from request form
	repo := r.FormValue("repo")
	branch := r.FormValue("branch")
	fileName := r.FormValue("fileName")

	// Extract file from request form
	file, header, err := r.FormFile(fileName)
	if err != nil {
		w.WriteHeader(http2.Status(http2.InternalServerError))
		c.Logger.GalError(err)
	} else {
		name := strings.Split(header.Filename, ".")
		fmt.Printf("Received file: %s\n", name[0])
		defer file.Close()
	}

	// Read file contents through buffer
	buffer := bytes.Buffer{}
	if _, err := io.Copy(&buffer, file); err != nil {
		w.WriteHeader(http2.Status(http2.InternalServerError))
		c.Logger.GalError(err)
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
		w.WriteHeader(http2.Status(http2.InternalServerError))
		c.Logger.GalError(err)
	} else { // Success
		marsh.MarshalResponse(res, w)
	}
}

// CreateRepository TODO
//
//
func (c *Config) CreateRepository(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Unmarshal create repository request
	req := &models.ReqCreateRepo{}
	marsh.UnmarshalRequest(req, w, r)

	// Create repository
	repo := github.Repository{
		Name:          req.RepoName,
		DefaultBranch: utils.String(consts.MASTER),
	}
	if res, _, err := c.GAL.Repositories.Create(ctx, consts.Z3E2C, &repo); err != nil {
		w.WriteHeader(http2.Status(http2.InternalServerError))
		c.Logger.GalError(err)
	} else { // Success
		marsh.MarshalResponse(res, w)
	}
}

// CreateBranch TODO
//
//
func (c *Config) CreateBranch(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Unmarshal create reference request
	req := &models.ReqCreateBranch{}
	marsh.UnmarshalRequest(req, w, r)

	// Get MASTER reference
	masterBranch, res, err := c.GAL.Git.GetRef(ctx, consts.Z3E2C, *req.RepoName, fmt.Sprintf("refs/heads/%s", consts.MASTER))
	if err != nil {
		w.WriteHeader(http2.Status(http2.InternalServerError))
		c.Logger.GalError(err)
	} else {
		fmt.Printf("Got master branch: %v\n", res)
	}

	// Create branch
	reference := github.Reference{
		Ref: utils.String(fmt.Sprintf("refs/heads/%s", *req.BranchName)),
		URL: utils.String(fmt.Sprintf("https://api.github.com/repos/%s/%s/git/refs/heads/%s", consts.Z3E2C, *req.RepoName, *req.BranchName)),
		Object: &github.GitObject{
			Type: utils.String("commit"),
			SHA:  masterBranch.Object.SHA,
			URL:  utils.String(fmt.Sprintf("https://api.github.com/repos/%s/%s/git/commits/%s", consts.Z3E2C, *req.RepoName, consts.MASTER)),
		},
	}
	if res, _, err := c.GAL.Git.CreateRef(ctx, consts.Z3E2C, *req.RepoName, &reference); err != nil {
		w.WriteHeader(http2.Status(http2.InternalServerError))
		c.Logger.GalError(err)
	} else { // Success
		marsh.MarshalResponse(res, w)
	}
}
