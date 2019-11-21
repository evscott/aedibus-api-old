package handlers

import (
	"context"
	"net/http"

	"github.com/evscott/z3-e2c-api/models"
	consts "github.com/evscott/z3-e2c-api/shared/constants"
	status "github.com/evscott/z3-e2c-api/shared/http-codes"
	"github.com/evscott/z3-e2c-api/shared/marsh"
)

//  TODO
//
//
func (c *Config) CreateComment(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	req := &models.ReqCreateComment{}
	marsh.UnmarshalRequest(req, w, r)

	c.helpers.CreateComment(ctx, w, *req.Path, *req.Body, *req.CommitID, *req.RepoName, *req.Position)

	w.WriteHeader(status.Status(status.OK))
}

// TODO
//
//
func (c *Config) UpdateAssignment(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	repoName := r.FormValue("repoName")
	branchName := r.FormValue("branchName")
	fileName := r.FormValue("fileName")
	contents := c.helpers.GetFileContents(w, r, fileName)

	c.helpers.UpdateFile(ctx, w, repoName, branchName, fileName, contents)
	c.helpers.UpdateAssignment(ctx, w, repoName, branchName)

	w.WriteHeader(status.Status(status.OK))
}

// TODO
//
//abc6
func (c *Config) UploadAssignment(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	repoName := r.FormValue("repoName")
	branchName := r.FormValue("branchName")
	fileName := r.FormValue("fileName")
	contents := c.helpers.GetFileContents(w, r, fileName)

	c.helpers.CreateFile(ctx, w, repoName, branchName, fileName, contents)
	c.helpers.UpdateAssignment(ctx, w, repoName, branchName)

	w.WriteHeader(status.Status(status.OK))
}

// TODO
//
//
func (c *Config) CreateAssignment(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	req := &models.ReqCreateRepo{}
	marsh.UnmarshalRequest(req, w, r)

	c.helpers.CreateRepository(ctx, w, *req.RepoName)
	c.helpers.CreateAssignment(ctx, w, *req.RepoName, consts.MASTER)

	w.WriteHeader(status.Status(status.OK))
}
