package handlers

import (
	"context"
	"net/http"

	"github.com/evscott/z3-e2c-api/models"
	consts "github.com/evscott/z3-e2c-api/shared/constants"
	"github.com/evscott/z3-e2c-api/shared/marsh"
)

// TODO
//
//
func (c *Config) CreateAssignment(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	req := &models.ReqCreateRepo{}
	marsh.UnmarshalRequest(req, w, r)

	c.helpers.CreateRepositoryHelper(ctx, w, *req.RepoName)
	c.helpers.CreateAssignmentHelper(ctx, w, *req.RepoName, consts.MASTER)
}

// TODO
//
//
func (c *Config) CreateAssignmentFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	repoName := r.FormValue("repoName")
	branchName := r.FormValue("branchName")
	fileName := r.FormValue("fileName")
	contents := c.helpers.ReceiveFileContentsHelper(w, r, fileName)

	c.helpers.CreateGitFile(ctx, w, repoName, branchName, fileName, contents)
	c.helpers.UpdateAssignmentBlob(ctx, w, repoName, branchName)
	// TODO create db file
}

// TODO
//
//
func (c *Config) UpdateAssignmentFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	repoName := r.FormValue("repoName")
	branchName := r.FormValue("branchName")
	fileName := r.FormValue("fileName")
	contents := c.helpers.ReceiveFileContentsHelper(w, r, fileName)

	c.helpers.UpdateFileHelper(ctx, w, repoName, branchName, fileName, contents)
	c.helpers.UpdateAssignmentBlob(ctx, w, repoName, branchName)
	// TODO update db file
}

// TODO
//
//
func (c *Config) GetFileContents(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	req := &models.ReqGetFile{}
	marsh.UnmarshalRequest(req, w, r)

	c.helpers.GetFileContentsHelper(ctx, w, *req.Name, *req.Branch, *req.Path)
}

//  TODO
//
//
func (c *Config) CreateComment(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	req := &models.ReqCreateComment{}
	marsh.UnmarshalRequest(req, w, r)

	c.helpers.CreateCommentHelper(ctx, w, *req.Path, *req.Body, *req.CommitID, *req.RepoName, *req.Position)
}
