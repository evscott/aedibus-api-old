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
func (c *Config) CreateSubmission(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Unmarshal create reference request
	req := &models.ReqCreateSubmission{}
	marsh.UnmarshalRequest(req, w, r)

	if err := c.helpers.GH.CreateGitBranch(ctx, req.Name, req.AssignmentName); err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.DB.CreateDbSubmission(ctx, req.Name, req.AssignmentName); err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	w.WriteHeader(status.Status(status.OK))
}

// TODO
//
//
func (c *Config) CreateSubmissionFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	assignmentName := r.FormValue("assignmentName")
	submissionName := r.FormValue("submissionName")
	fileName := r.FormValue("fileName")

	contents, err := c.helpers.DB.ReceiveFileContentsHelper(w, r, fileName)
	if err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.GH.CreateGitFile(ctx, assignmentName, submissionName, fileName, contents); err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	blobSHA, err := c.helpers.GH.GetMasterBlobSha(ctx, assignmentName)
	if err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.DB.UpdateAssignmentBlob(ctx, assignmentName, *blobSHA); err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	submission, err := c.helpers.DB.GetSubmissionByBranchAndRepo(ctx, assignmentName, submissionName)
	if err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.DB.CreateDbFile(ctx, fileName, submission.Name); err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	w.WriteHeader(status.Status(status.OK))
}
