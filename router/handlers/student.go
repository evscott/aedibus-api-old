package handlers

import (
	"context"
	"net/http"

	"github.com/evscott/z3-e2c-api/models"
	status "github.com/evscott/z3-e2c-api/shared/http-codes"
	"github.com/evscott/z3-e2c-api/shared/marsh"
)

// TODO
//
//
func (c *Config) SubmitAssignment(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	req := &models.ReqCreatePR{}
	marsh.UnmarshalRequest(req, w, r)

	c.helpers.CreatePullRequestHelper(ctx, w, *req.Title, *req.Head, *req.Body, *req.RepoName)
	w.WriteHeader(status.Status(status.OK))
}

// TODO
//
//
func (c *Config) CreateSubmission(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Unmarshal create reference request
	req := &models.ReqCreateBranch{}
	marsh.UnmarshalRequest(req, w, r)

	c.helpers.CreateBranchHelper(ctx, w, *req.RepoName, *req.BranchName)
	w.WriteHeader(status.Status(status.OK))
}
