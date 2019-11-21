package handlers

import (
	"context"
	"net/http"

	"github.com/evscott/z3-e2c-api/models"
	consts "github.com/evscott/z3-e2c-api/shared/constants"
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
}

// TODO
//
//
func (c *Config) UpdateFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	repoName := r.FormValue("repoName")
	branchName := r.FormValue("branchName")
	fileName := r.FormValue("fileName")
	contents := c.helpers.ReceiveFileContents(w, r, fileName)

	c.helpers.UpdateFile(ctx, w, repoName, branchName, fileName, contents)
	c.helpers.UpdateAssignment(ctx, w, repoName, branchName)
}

// TODO
//
//abc6
func (c *Config) UploadFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	repoName := r.FormValue("repoName")
	branchName := r.FormValue("branchName")
	fileName := r.FormValue("fileName")
	contents := c.helpers.ReceiveFileContents(w, r, fileName)

	c.helpers.CreateFile(ctx, w, repoName, branchName, fileName, contents)
	c.helpers.UpdateAssignment(ctx, w, repoName, branchName)
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
}

// TODO
//
//
func (c *Config) GetReadme(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	req := &models.ReqGetAssignment{}
	marsh.UnmarshalRequest(req, w, r)

	c.helpers.GetReadme(ctx, w, *req.Name, *req.Branch)
}
