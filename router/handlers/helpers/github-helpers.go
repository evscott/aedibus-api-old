package helpers

import (
	"context"
	"fmt"
	"net/http"

	consts "github.com/evscott/z3-e2c-api/shared/constants"
	status "github.com/evscott/z3-e2c-api/shared/http-codes"
	"github.com/evscott/z3-e2c-api/shared/marsh"
	"github.com/evscott/z3-e2c-api/shared/utils"
	"github.com/google/go-github/github"
)

// TODO
//
func (c *Config) CreatePullRequest(ctx context.Context, w http.ResponseWriter, title, head, body, repoName string) {
	pullRequest := github.NewPullRequest{
		Title:               &title,
		Head:                &head,
		Base:                utils.String(consts.MASTER),
		Body:                &body,
		Issue:               nil,
		MaintainerCanModify: utils.Bool(true),
	}
	if res, _, err := c.GAL.PullRequests.Create(ctx, consts.Z3E2C, repoName, &pullRequest); err != nil {
		w.WriteHeader(status.Status(status.InternalServerError))
		c.Logger.GalError(err)
	} else { // Success
		marsh.MarshalResponse(res, w)
	}
}

// TODO
//
func (c *Config) CreateRepository(ctx context.Context, w http.ResponseWriter, repoName string) {
	repo := github.Repository{
		Name:          &repoName,
		DefaultBranch: utils.String(consts.MASTER),
	}
	if _, _, err := c.GAL.Repositories.Create(ctx, consts.Z3E2C, &repo); err != nil {
		w.WriteHeader(status.Status(status.InternalServerError))
		c.Logger.GalError(err)
	}
}

// TODO
//
func (c *Config) CreateFile(ctx context.Context, w http.ResponseWriter, repoName, branchName, fileName string, contents []byte) {
	fileOptions := github.RepositoryContentFileOptions{
		Message: utils.String("Uploading file"),
		Content: contents,
		Branch:  &branchName,
	}
	if _, _, err := c.GAL.Repositories.CreateFile(ctx, consts.Z3E2C, repoName, fileName, &fileOptions); err != nil {
		w.WriteHeader(status.Status(status.InternalServerError))
		c.Logger.GalError(err)
	}
}

// TODO
//
func (c *Config) UpdateFile(ctx context.Context, w http.ResponseWriter, repo, branch, fileName string, contents []byte) {
	// Get blob sha of file from Github to be used as target of update
	var sha string
	getOptions := github.RepositoryContentGetOptions{Ref: fmt.Sprintf("heads/%s", branch)}
	if contents, _, _, err := c.GAL.Repositories.GetContents(ctx, consts.Z3E2C, repo, fileName, &getOptions); err != nil {
		w.WriteHeader(status.Status(status.InternalServerError))
		c.Logger.GalError(err)
	} else {
		sha = *contents.SHA
	}

	// Upload file to Github
	fileOptions := github.RepositoryContentFileOptions{
		Message: utils.String("Updating file"), // TODO
		Content: contents,
		Branch:  &branch,
		SHA:     &sha,
	}
	if _, _, err := c.GAL.Repositories.UpdateFile(ctx, consts.Z3E2C, repo, fileName, &fileOptions); err != nil {
		w.WriteHeader(status.Status(status.InternalServerError))
		c.Logger.GalError(err)
	}
}

// TODO
//
func (c *Config) CreateBranch(ctx context.Context, w http.ResponseWriter, repoName, branchName string) {
	masterBranch, _, err := c.GAL.Git.GetRef(ctx, consts.Z3E2C, repoName, fmt.Sprintf("refs/heads/%s", consts.MASTER))
	if err != nil {
		w.WriteHeader(status.Status(status.InternalServerError))
		c.Logger.GalError(err)
	}

	reference := github.Reference{
		Ref: utils.String(fmt.Sprintf("refs/heads/%s", branchName)),
		URL: utils.String(fmt.Sprintf("https://api.github.com/repos/%s/%s/git/refs/heads/%s", consts.Z3E2C, repoName, branchName)),
		Object: &github.GitObject{
			Type: utils.String("commit"),
			SHA:  masterBranch.Object.SHA,
			URL:  utils.String(fmt.Sprintf("https://api.github.com/repos/%s/%s/git/commits/%s", consts.Z3E2C, repoName, consts.MASTER)),
		},
	}

	if _, _, err := c.GAL.Git.CreateRef(ctx, consts.Z3E2C, repoName, &reference); err != nil {
		w.WriteHeader(status.Status(status.InternalServerError))
		c.Logger.GalError(err)
	}
}
