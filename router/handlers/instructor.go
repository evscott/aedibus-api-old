package handlers

import (
	"context"
	status "github.com/evscott/z3-e2c-api/shared/http-codes"
	"net/http"

	"github.com/evscott/z3-e2c-api/models"
	"github.com/evscott/z3-e2c-api/shared/marsh"
)

// TODO
//
//
func (c *Config) CreateAssignment(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	req := &models.ReqCreateAssignment{}
	marsh.UnmarshalRequest(req, w, r)

	if err := c.helpers.GH.CreateRepository(ctx, req.Name); err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}
	if err := c.helpers.DB.CreateAssignment(ctx, req.Name); err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}
	w.WriteHeader(status.Status(status.OK))
}

// TODOName
//
//
func (c *Config) CreateAssignmentFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	assignmentName := r.FormValue("assignmentName")
	submissionName := r.FormValue("submissionName")
	fileName := r.FormValue("fileName")

	contents, err := c.helpers.DB.GetFileFromForm(w, r, fileName)
	if err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.GH.CreateFile(ctx, assignmentName, submissionName, fileName, contents); err != nil {
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

	w.WriteHeader(status.Status(status.OK))
}

// TODO
//
//
func (c *Config) UpdateAssignmentFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	assignmentName := r.FormValue("assignmentName")
	submissionName := r.FormValue("submissionName")
	fileName := r.FormValue("fileName")

	contents, err := c.helpers.DB.GetFileFromForm(w, r, fileName)
	if err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.GH.UpdateFile(ctx, assignmentName, submissionName, fileName, contents); err != nil {
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

	w.WriteHeader(status.Status(status.OK))
}

// TODO
//
//
func (c *Config) GetFileContents(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	req := &models.ReqGetFile{}
	marsh.UnmarshalRequest(req, w, r)

	res, err := c.helpers.GH.GetFileContents(ctx, req.FileName, req.SubmissionName)
	if err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	marsh.MarshalResponse(res, w)
}
