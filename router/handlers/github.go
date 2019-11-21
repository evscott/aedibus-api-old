package handlers

import (
	"bytes"
	"context"
	"fmt"
	"github.com/evscott/z3-e2c-api/dal"
	status "github.com/evscott/z3-e2c-api/shared/http-codes"
	"github.com/evscott/z3-e2c-api/shared/marsh"
	"io"
	"net/http"

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

	req := &models.ReqCreateComment{}
	marsh.UnmarshalRequest(req, w, r)

	c.createComment(ctx, w, *req.Path, *req.Body, *req.CommitID, *req.RepoName, *req.Position)
	w.WriteHeader(status.Status(status.OK))
}

// CreatePullRequest TODO
//
//
func (c *Config) CreatePullRequest(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	req := &models.ReqCreatePR{}
	marsh.UnmarshalRequest(req, w, r)

	c.createPullRequest(ctx, w, *req.Title, *req.Head, *req.Body, *req.RepoName)
	w.WriteHeader(status.Status(status.OK))
}

// UpdateFile
//
//
func (c *Config) UpdateFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	repo := r.FormValue("repo")
	branch := r.FormValue("branch")
	fileName := r.FormValue("fileName")
	contents := c.getFileContents(w, r, fileName)

	c.updateFile(ctx, w, repo, branch, fileName, contents)
	c.updateAssignment(ctx, w, repo)

	w.WriteHeader(status.Status(status.OK))
}

// UploadFile TODO
//
//
func (c *Config) UploadFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	repoName := r.FormValue("repo")
	branchName := r.FormValue("branch")
	fileName := r.FormValue("fileName")
	contents := c.getFileContents(w, r, fileName)

	c.createGithubFile(ctx, w, repoName, branchName, fileName, contents)
	c.updateAssignment(ctx, w, repoName)

	w.WriteHeader(status.Status(status.OK))
}

// CreateRepository TODO
//
//
func (c *Config) CreateRepository(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	req := &models.ReqCreateRepo{}
	marsh.UnmarshalRequest(req, w, r)

	c.createRepository(ctx, w, *req.RepoName)
	c.createAssignment(ctx, w, *req.RepoName, consts.MASTER)

	w.WriteHeader(status.Status(status.OK))
}

// CreateBranch TODO
//
//
func (c *Config) CreateBranch(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Unmarshal create reference request
	req := &models.ReqCreateBranch{}
	marsh.UnmarshalRequest(req, w, r)

	c.createBranch(ctx, w, *req.RepoName, *req.BranchName)
}

// _    _   ______   _        _____    ______   _____
// | |  | | |  ____| | |      |  __ \  |  ____| |  __ \
// | |__| | | |__    | |      | |__) | | |__    | |__) |
// |  __  | |  __|   | |      |  ___/  |  __|   |  _  /
// | |  | | | |____  | |____  | |      | |____  | | \ \
// |_|  |_| |______| |______| |_|      |______| |_|  \_\
//

func (c *Config) createComment(ctx context.Context, w http.ResponseWriter, path, body, commitID, repoName string, position int) {
	comment := github.PullRequestComment{
		Path:     &path,
		Body:     &body,
		Position: &position,
		CommitID: &commitID,
	}
	if res, _, err := c.GAL.PullRequests.CreateComment(ctx, consts.Z3E2C, repoName, 1, &comment); err != nil {
		w.WriteHeader(status.Status(status.InternalServerError))
		c.Logger.GalError(err)
	} else { // Success
		marsh.MarshalResponse(res, w)
	}
}

func (c *Config) createPullRequest(ctx context.Context, w http.ResponseWriter, title, head, body, repoName string) {
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

func (c *Config) createRepository(ctx context.Context, w http.ResponseWriter, repoName string) {
	repo := github.Repository{
		Name:          &repoName,
		DefaultBranch: utils.String(consts.MASTER),
	}
	if _, _, err := c.GAL.Repositories.Create(ctx, consts.Z3E2C, &repo); err != nil {
		w.WriteHeader(status.Status(status.InternalServerError))
		c.Logger.GalError(err)
	}
}

func (c *Config) createAssignment(ctx context.Context, w http.ResponseWriter, repoName, branchName string) {
	assignment := &models.Assignment{
		Name:   repoName,
		Branch: branchName,
	}
	if err := c.DAL.Provider.CreateAssignment(ctx, assignment); err != nil {
		w.WriteHeader(status.Status(status.InternalServerError))
		c.Logger.DalError(err)
	} else {
		w.WriteHeader(status.Status(status.OK))
	}
}

func (c *Config) getFileContents(w http.ResponseWriter, r *http.Request, fileName string) []byte {
	file, _, err := r.FormFile(fileName)
	if err != nil {
		w.WriteHeader(status.Status(status.InternalServerError))
		c.Logger.GalError(err)
	} else {
		defer file.Close()
	}
	// Read file contents through buffer
	buffer := bytes.Buffer{}
	if _, err := io.Copy(&buffer, file); err != nil {
		w.WriteHeader(status.Status(status.InternalServerError))
		c.Logger.GalError(err)
	}

	defer buffer.Reset()
	return buffer.Bytes()
}

func (c *Config) createGithubFile(ctx context.Context, w http.ResponseWriter, repoName, branchName, fileName string, contents []byte) {
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

func (c *Config) updateFile(ctx context.Context, w http.ResponseWriter, repo, branch, fileName string, contents []byte) {
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
} // Create repository

func (c *Config) updateAssignment(ctx context.Context, w http.ResponseWriter, repo string) {
	masterBranch, _, err := c.GAL.Git.GetRef(ctx, consts.Z3E2C, repo, fmt.Sprintf("refs/heads/%s", consts.MASTER))
	if err != nil {
		w.WriteHeader(status.Status(status.InternalServerError))
		c.Logger.GalError(err)
	}
	assignment := &models.Assignment{
		Name:     repo,
		BlobShah: *masterBranch.Object.SHA,
	}
	if err := c.DAL.Provider.UpdateAssignment(ctx, assignment); err != nil {
		w.WriteHeader(status.Status(status.InternalServerError))
	}
}

func (c *Config) createBranch(ctx context.Context, w http.ResponseWriter, repoName, branchName string) {
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
