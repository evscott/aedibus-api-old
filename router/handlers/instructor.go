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
	if err := marsh.UnmarshalRequest(req, w, r); err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.GH.CreateRepository(ctx, req.Name); err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.DB.CreateAssignment(ctx, req.Name); err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}
}

// TODOName
//
//
func (c *Config) CreateAssignmentFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	assignmentName := r.FormValue("assignmentName")
	dropboxName := r.FormValue("dropboxName")
	fileName := r.FormValue("fileName")

	contents, err := c.helpers.DB.GetFileFromForm(w, r, fileName)
	if err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.GH.CreateFile(ctx, assignmentName, dropboxName, fileName, contents); err != nil {
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
}

// TODO
//
//
func (c *Config) UpdateAssignmentFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	assignmentName := r.FormValue("assignmentName")
	dropboxName := r.FormValue("dropboxName")
	fileName := r.FormValue("fileName")

	contents, err := c.helpers.DB.GetFileFromForm(w, r, fileName)
	if err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.GH.UpdateFile(ctx, assignmentName, dropboxName, fileName, contents); err != nil {
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
}

// TODO
//
//
func (c *Config) GetFileContents(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	req := &models.ReqGetFile{}
	if err := marsh.UnmarshalRequest(req, w, r); err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	res, err := c.helpers.GH.GetFileContents(ctx, req.FileName, req.DropboxName)
	if err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := marsh.MarshalResponse(res, w); err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}
}

// TODO
//
//
func (c *Config) CreateDropbox(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Unmarshal create reference request
	req := &models.ReqCreateDropbox{}
	if err := marsh.UnmarshalRequest(req, w, r); err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.GH.CreateDropbox(ctx, req.DropboxName, req.AssignmentName); err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.DB.CreateDropbox(ctx, req.DropboxName, req.AssignmentName); err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}
}
