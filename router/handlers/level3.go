package handlers

import (
	"context"
	consts "github.com/evscott/aedibus-api/shared/constants"
	status "github.com/evscott/aedibus-api/shared/http-codes"
	"github.com/evscott/aedibus-api/shared/utils"
	"net/http"

	"github.com/evscott/aedibus-api/models"
	"github.com/evscott/aedibus-api/shared/marsh"
)

// TODO
//
//
func (c *Config) CreateAssignment(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	req := &models.ReqCreateAssignment{}
	if err := marsh.UnmarshalRequest(req, w, r); err != nil {
		c.logger.MarshError("unmarshalling request", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.GH.CreateRepository(ctx, req.AssignmentName); err != nil {
		c.logger.GalError("creating assignment (repository) on Github", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	assignment, err := c.helpers.DB.CreateAssignment(ctx, req.AssignmentName)
	if err != nil {
		c.logger.DalError("creating assignment in database", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	dropbox, err := c.helpers.DB.CreateDropbox(ctx, consts.MASTER, assignment.ID)
	if err != nil {
		c.logger.DalError("creating dropbox in database", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	res, err := c.helpers.GH.CreateFile(ctx, assignment.Name, consts.MASTER, consts.README, []byte(req.ReadmeContents))
	if err != nil {
		c.logger.DalError("creating readme file on Github", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.DB.CreateFile(ctx, consts.README, assignment.ID, dropbox.ID, *res.Commit.SHA); err != nil {
		c.logger.DalError("creating readme file in database", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	blobSHA, err := c.helpers.GH.GetMasterBlobSha(ctx, assignment.Name)
	if err != nil {
		c.logger.GalError("getting blob sha from Github", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.DB.UpdateAssignmentBlob(ctx, assignment.ID, *blobSHA); err != nil {
		c.logger.DalError("updating blob sha in database", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	w.WriteHeader(status.Status(status.OK))
}

// TODOs
//
//
func (c *Config) CreateAssignmentFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	assignmentName := r.FormValue("assignmentName")
	dropboxName := r.FormValue("dropboxName")
	fileName := r.FormValue("fileName")

	contents, err := utils.GetFileFromForm(r, fileName)
	if err != nil {
		c.logger.UtilsError("getting file from form", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	res, err := c.helpers.GH.CreateFile(ctx, assignmentName, dropboxName, fileName, contents)
	if err != nil {
		c.logger.GalError("creating file", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.DB.CreateFile(ctx, fileName, assignmentName, consts.MASTER, *res.Commit.SHA); err != nil {
		c.logger.DalError("creating file", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	blobSHA, err := c.helpers.GH.GetMasterBlobSha(ctx, assignmentName)
	if err != nil {
		c.logger.GalError("getting blob sha", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.DB.UpdateAssignmentBlob(ctx, assignmentName, *blobSHA); err != nil {
		c.logger.GalError("updating blob sha", err)
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
	if err := marsh.UnmarshalRequest(req, w, r); err != nil {
		c.logger.MarshError("unmarshalling request", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	res, err := c.helpers.GH.GetFileContents(ctx, req.FileName, req.AssignmentName)
	if err != nil {
		c.logger.GalError("getting file contents", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := marsh.MarshalResponse(res, w); err != nil {
		c.logger.MarshError("marshalling response", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	w.WriteHeader(status.Status(status.OK))
}

// TODO
//
//
func (c *Config) CreateDropbox(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Unmarshal create reference request
	req := &models.ReqCreateDropbox{}
	if err := marsh.UnmarshalRequest(req, w, r); err != nil {
		c.logger.MarshError("unmarshalling response", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.GH.CreateDropbox(ctx, req.DropboxName, req.AID); err != nil {
		c.logger.GalError("creating dropbox", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if _, err := c.helpers.DB.CreateDropbox(ctx, req.DropboxName, req.AID); err != nil {
		c.logger.DalError("creating dropbox", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	w.WriteHeader(status.Status(status.OK))
}
