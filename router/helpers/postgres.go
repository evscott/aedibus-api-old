package helpers

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/evscott/z3-e2c-api/models"
	consts "github.com/evscott/z3-e2c-api/shared/constants"
	status "github.com/evscott/z3-e2c-api/shared/http-codes"
	"github.com/evscott/z3-e2c-api/shared/marsh"
	"github.com/google/go-github/github"
)

// TODO
//
func (c *Config) CreateCommentHelper(ctx context.Context, w http.ResponseWriter, path, body, commitID, repoName string, position int) {
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

// TODO
//
func (c *Config) CreateAssignmentHelper(ctx context.Context, w http.ResponseWriter, repoName, branchName string) {
	assignment := &models.Assignment{
		Name:   &repoName,
		Branch: &branchName,
	}
	if err := c.DAL.Provider.CreateAssignment(ctx, assignment); err != nil {
		w.WriteHeader(status.Status(status.InternalServerError))
		c.Logger.DalError(err)
	} else {
		w.WriteHeader(status.Status(status.OK))
	}
}

// TODO
//
func (c *Config) ReceiveFileContentsHelper(w http.ResponseWriter, r *http.Request, fileName string) []byte {
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

// TODO
//
func (c *Config) UpdateAssignmentHelper(ctx context.Context, w http.ResponseWriter, repo, branch, fileName string) {
	masterBranch, _, err := c.GAL.Git.GetRef(ctx, consts.Z3E2C, repo, fmt.Sprintf("refs/heads/%s", consts.MASTER))
	if err != nil {
		w.WriteHeader(status.Status(status.InternalServerError))
		c.Logger.GalError(err)
	}
	assignment := &models.Assignment{
		Name:     &repo,
		Branch:   &branch,
		BlobShah: masterBranch.Object.SHA,
	}
	if err := c.DAL.Provider.UpdateAssignment(ctx, assignment); err != nil {
		w.WriteHeader(status.Status(status.InternalServerError))
	}
}
